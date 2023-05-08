package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"cosmossdk.io/math"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	codectypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	"github.com/cosmos/cosmos-sdk/x/authz"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	crisistypes "github.com/cosmos/cosmos-sdk/x/crisis/types"
	distributiontypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	evidencetypes "github.com/cosmos/cosmos-sdk/x/evidence/types"
	feegranttypes "github.com/cosmos/cosmos-sdk/x/feegrant"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	govtypes "github.com/cosmos/cosmos-sdk/x/gov/types"
	govv1types "github.com/cosmos/cosmos-sdk/x/gov/types/v1"
	"github.com/cosmos/cosmos-sdk/x/group"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	ibchost "github.com/cosmos/ibc-go/v5/modules/core/24-host"
	ibctypes "github.com/cosmos/ibc-go/v5/modules/core/types"
	"github.com/gitopia/gitopia/app/params"
	"github.com/gitopia/gitopia/x/gitopia/keeper"
	testnettypes "github.com/gitopia/gitopia/x/gitopia/migrations/testnet"
	gitopiatypes "github.com/gitopia/gitopia/x/gitopia/types"
	rewardstypes "github.com/gitopia/gitopia/x/rewards/types"
	"github.com/spf13/cobra"
	tmjson "github.com/tendermint/tendermint/libs/json"
	tmtypes "github.com/tendermint/tendermint/types"
)

const (
	flagGenesisTime   = "genesis-time"
	flagInitialHeight = "initial-height"
)

const (
	rewardsServiceAddress            = "gitopia1a875smmd9va45tsx398prdzjtm5fg23mlzzgck"
	strategicReserveAddress1         = ""
	strategicReserveAddress2         = ""
	strategicReserveAddress3         = ""
	strategicReserveAddress4         = ""
	strategicReserveAddress5         = ""
	feegrantsAddress                 = ""
	earlySupportersMultiSigAddress   = ""
	strategicPartnersMultiSigAddress = ""
	advisorsMultiSigAddress          = ""
	teamMultiSigAddress              = ""
)

const (
	STRATEGIC_PARTNERS_AMOUNT = 51_071_429_000_000
	ADVISORS_AMOUNT           = 10_000_000_000_000
)

func normalizeRepoName(name string) string {
	return strings.ToLower(name)
}

func newAnyAuthorization(a authz.Authorization) *codectypes.Any {
	any, err := codectypes.NewAnyWithValue(a)
	if err != nil {
		panic(err)
	}

	return any
}

func resolveRepoNameConflict(name string, repoNameMap map[string]int) string {
	normalizedRepoName := normalizeRepoName(name)
	count := repoNameMap[normalizedRepoName]
	if count > 0 {
		name = fmt.Sprintf("%s-%d", name, count)
	}
	repoNameMap[normalizedRepoName]++
	return name
}

// migrateTestnetState performs type migrations from v1.2.0 to v1.4.0. The
// migration includes:
//
// - Removed `issues` and `pullRequests` from Repository.
// - Add `repositoryId` in Comment and Bounty.
// - Modified comment structure - Parent: issue and pull; various comment types like label, assignees etc; reactions; replies; resolved/unresolved
func migrateTestnetState(state testnettypes.GenesisState) (gitopiatypes.GenesisState, error) {
	var genesisState gitopiatypes.GenesisState

	genesisState.Params = gitopiatypes.Params{
		NextInflationTime: time.Now().AddDate(2, 0, 0),
		PoolProportions: gitopiatypes.PoolProportions{
			Ecosystem: &gitopiatypes.DistributionProportion{Proportion: sdk.MustNewDecFromStr("30.0")},
			Team:      &gitopiatypes.DistributionProportion{Proportion: sdk.MustNewDecFromStr("28.0")},
		},
		TeamProportions: []gitopiatypes.DistributionProportion{
			{Proportion: sdk.MustNewDecFromStr("35.0"), Address: ""},
			{Proportion: sdk.MustNewDecFromStr("35.0"), Address: "gitopia14t0ta8vvv2nrcx86g87z888s7pqat4svuyw7ae"},
			{Proportion: sdk.MustNewDecFromStr("12.5"), Address: "gitopia1gyldx4ysv8u97v7rnjuw06sq35d8khmvn28d9n"},
			{Proportion: sdk.MustNewDecFromStr("2.0"), Address: "gitopia1ps5vrjmhtrkyxrge7d0fwvzsf02lq49wq3xeau"},
			{Proportion: sdk.MustNewDecFromStr("2.0"), Address: "gitopia1g0nvcrrd59zef2r9jt56jvut3gf6040svuveaa"},
			{Proportion: sdk.MustNewDecFromStr("2.0"), Address: "gitopia1kcnhjh9fkcc2w74g20s8edu6ue2eyamd3aqees"},
			{Proportion: sdk.MustNewDecFromStr("2.0"), Address: ""},
			{Proportion: sdk.MustNewDecFromStr("1.0"), Address: ""},
			{Proportion: sdk.MustNewDecFromStr("1.0"), Address: "gitopia1fyztge5vhfdfhnghumssv7z5cca7ctyecq3fqf"},
			// 6.5 + 1.0
			{Proportion: sdk.MustNewDecFromStr("7.5"), Address: teamMultiSigAddress},
		},
		GitServer:       "gitopia1a875smmd9va45tsx398prdzjtm5fg23mlzzgck",
		StorageProvider: "gitopia1a875smmd9va45tsx398prdzjtm5fg23mlzzgck",
	}

	for _, oldTask := range state.TaskList {
		task := gitopiatypes.Task{
			Id:       oldTask.Id,
			Type:     gitopiatypes.TaskType(oldTask.Type),
			State:    gitopiatypes.TaskState(oldTask.State),
			Message:  oldTask.Message,
			Creator:  oldTask.Creator,
			Provider: oldTask.Provider,
		}

		genesisState.TaskList = append(genesisState.TaskList, task)
	}

	genesisState.TaskCount = state.TaskCount

	for _, oldBranch := range state.BranchList {
		branch := gitopiatypes.Branch{
			Id:             oldBranch.Id,
			RepositoryId:   oldBranch.RepositoryId,
			Name:           oldBranch.Name,
			Sha:            oldBranch.Sha,
			AllowForcePush: oldBranch.AllowForcePush,
			CreatedAt:      oldBranch.CreatedAt,
			UpdatedAt:      oldBranch.UpdatedAt,
		}

		genesisState.BranchList = append(genesisState.BranchList, branch)
	}

	genesisState.BranchCount = state.BranchCount

	for _, oldTag := range state.TagList {
		tag := gitopiatypes.Tag{
			Id:           oldTag.Id,
			RepositoryId: oldTag.RepositoryId,
			Name:         oldTag.Name,
			Sha:          oldTag.Sha,
			CreatedAt:    oldTag.CreatedAt,
			UpdatedAt:    oldTag.UpdatedAt,
		}

		genesisState.TagList = append(genesisState.TagList, tag)
	}

	genesisState.TagCount = state.TagCount

	for _, oldMember := range state.MemberList {
		member := gitopiatypes.Member{
			Id:         oldMember.Id,
			Address:    oldMember.Address,
			DaoAddress: oldMember.DaoAddress,
			Role:       gitopiatypes.MemberRole(oldMember.Role),
		}

		genesisState.MemberList = append(genesisState.MemberList, member)
	}

	genesisState.MemberCount = state.MemberCount

	for _, oldRelease := range state.ReleaseList {
		var attachments []*gitopiatypes.Attachment
		for _, oldAttachment := range oldRelease.Attachments {
			attachments = append(attachments, &gitopiatypes.Attachment{
				Name:     oldAttachment.Name,
				Size_:    oldAttachment.Size_,
				Sha:      oldAttachment.Sha,
				Uploader: oldAttachment.Uploader,
			})
		}

		release := gitopiatypes.Release{
			Creator:      oldRelease.Creator,
			Id:           oldRelease.Id,
			RepositoryId: oldRelease.RepositoryId,
			TagName:      oldRelease.TagName,
			Target:       oldRelease.Target,
			Name:         oldRelease.Name,
			Description:  oldRelease.Description,
			Attachments:  attachments,
			Draft:        oldRelease.Draft,
			PreRelease:   oldRelease.PreRelease,
			IsTag:        oldRelease.IsTag,
			CreatedAt:    oldRelease.CreatedAt,
			UpdatedAt:    oldRelease.UpdatedAt,
			PublishedAt:  oldRelease.PublishedAt,
		}

		genesisState.ReleaseList = append(genesisState.ReleaseList, release)
	}

	genesisState.ReleaseCount = state.ReleaseCount

	for _, oldDao := range state.DaoList {
		dao := gitopiatypes.Dao{
			Creator:     oldDao.Creator,
			Id:          oldDao.Id,
			Address:     oldDao.Address,
			Name:        oldDao.Name,
			AvatarUrl:   oldDao.AvatarUrl,
			Followers:   oldDao.Followers,
			Following:   oldDao.Following,
			Teams:       oldDao.Teams,
			Location:    oldDao.Location,
			Website:     oldDao.Website,
			Verified:    oldDao.Verified,
			Description: oldDao.Description,
			CreatedAt:   oldDao.CreatedAt,
			UpdatedAt:   oldDao.UpdatedAt,
		}

		genesisState.DaoList = append(genesisState.DaoList, dao)
	}

	genesisState.DaoCount = state.DaoCount

	for _, oldUser := range state.UserList {
		if oldUser.Creator == "" {
			continue
		}
		user := gitopiatypes.User{
			Creator:        oldUser.Creator,
			Id:             oldUser.Id,
			Name:           oldUser.Name,
			Username:       oldUser.Username,
			UsernameGithub: oldUser.UsernameGithub,
			AvatarUrl:      oldUser.AvatarUrl,
			Followers:      oldUser.Followers,
			Following:      oldUser.Following,
			StarredRepos:   oldUser.StarredRepos,
			Subscriptions:  oldUser.Subscriptions,
			Bio:            oldUser.Bio,
			CreatedAt:      oldUser.CreatedAt,
			UpdatedAt:      oldUser.UpdatedAt,
			Verified:       oldUser.Verified,
		}

		genesisState.UserList = append(genesisState.UserList, user)
	}

	genesisState.UserCount = state.UserCount

	for _, oldUserDao := range state.UserDaoList {
		userDao := gitopiatypes.UserDao{
			UserAddress: oldUserDao.UserAddress,
			DaoAddress:  oldUserDao.DaoAddress,
		}

		genesisState.UserDaoList = append(genesisState.UserDaoList, userDao)
	}

	for _, oldWhois := range state.WhoisList {
		whois := gitopiatypes.Whois{
			Creator:   oldWhois.Creator,
			Id:        oldWhois.Id,
			Name:      oldWhois.Name,
			Address:   oldWhois.Address,
			OwnerType: gitopiatypes.OwnerType(oldWhois.OwnerType),
		}

		genesisState.WhoisList = append(genesisState.WhoisList, whois)
	}

	genesisState.WhoisCount = state.WhoisCount

	repoNameMap := make(map[string]map[string]int)
	newRepoNameMap := make(map[string]map[string]string)

	for _, oldRepository := range state.RepositoryList {
		var labels []*gitopiatypes.RepositoryLabel
		var releases []*gitopiatypes.RepositoryRelease
		var collaborators []*gitopiatypes.RepositoryCollaborator
		var backups []*gitopiatypes.RepositoryBackup

		for _, oldLabel := range oldRepository.Labels {
			labels = append(labels, &gitopiatypes.RepositoryLabel{
				Id:          oldLabel.Id,
				Name:        oldLabel.Name,
				Color:       oldLabel.Color,
				Description: oldLabel.Description,
			})
		}

		for _, oldRelease := range oldRepository.Releases {
			releases = append(releases, &gitopiatypes.RepositoryRelease{
				Id:      oldRelease.Id,
				TagName: oldRelease.TagName,
			})
		}

		for _, oldCollaborator := range oldRepository.Collaborators {
			collaborators = append(collaborators, &gitopiatypes.RepositoryCollaborator{
				Id:         oldCollaborator.Id,
				Permission: gitopiatypes.RepositoryCollaborator_Permission(oldCollaborator.Permission),
			})
		}

		for _, oldBackup := range oldRepository.Backups {
			backups = append(backups, &gitopiatypes.RepositoryBackup{
				Store: gitopiatypes.RepositoryBackup_Store(oldBackup.Store),
				Refs:  oldBackup.Refs,
			})
		}

		if _, ok := repoNameMap[oldRepository.Owner.Id]; !ok {
			repoNameMap[oldRepository.Owner.Id] = make(map[string]int)
			newRepoNameMap[oldRepository.Owner.Id] = make(map[string]string)
		}
		newRepoName := resolveRepoNameConflict(oldRepository.Name, repoNameMap[oldRepository.Owner.Id])
		newRepoNameMap[oldRepository.Owner.Id][oldRepository.Name] = newRepoName

		repository := gitopiatypes.Repository{
			Creator: oldRepository.Creator,
			Id:      oldRepository.Id,
			Name:    newRepoName,
			Owner: &gitopiatypes.RepositoryOwner{
				Id:   oldRepository.Owner.Id,
				Type: gitopiatypes.OwnerType(oldRepository.Owner.Type),
			},
			Description:         oldRepository.Description,
			Forks:               oldRepository.Forks,
			Subscribers:         oldRepository.Subscribers,
			Commits:             oldRepository.Commits,
			IssuesCount:         oldRepository.IssuesCount,
			PullsCount:          oldRepository.PullsCount,
			Labels:              labels,
			LabelsCount:         oldRepository.LabelsCount,
			Releases:            releases,
			CreatedAt:           oldRepository.CreatedAt,
			UpdatedAt:           oldRepository.UpdatedAt,
			PushedAt:            oldRepository.PushedAt,
			Stargazers:          oldRepository.Stargazers,
			Archived:            oldRepository.Archived,
			License:             oldRepository.License,
			DefaultBranch:       oldRepository.DefaultBranch,
			Parent:              oldRepository.Parent,
			Fork:                oldRepository.Fork,
			Collaborators:       collaborators,
			AllowForking:        oldRepository.AllowForking,
			Backups:             backups,
			EnableArweaveBackup: oldRepository.EnableArweaveBackup,
		}

		genesisState.RepositoryList = append(genesisState.RepositoryList, repository)
	}

	genesisState.RepositoryCount = state.RepositoryCount

	for _, oldbaseRepositoryKey := range state.BaseRepositoryKeyList {
		baseRepositoryKey := gitopiatypes.BaseRepositoryKey{
			Id:      oldbaseRepositoryKey.Id,
			Address: oldbaseRepositoryKey.Address,
			Name:    normalizeRepoName(newRepoNameMap[oldbaseRepositoryKey.Address][oldbaseRepositoryKey.Name]),
		}

		genesisState.BaseRepositoryKeyList = append(genesisState.BaseRepositoryKeyList, baseRepositoryKey)
	}

	commentMap := make(map[uint64]testnettypes.Comment)
	for _, comment := range state.CommentList {
		commentMap[comment.Id] = comment
	}

	genesisState.CommentCount = state.CommentCount

	for _, oldIssue := range state.IssueList {
		issue := gitopiatypes.Issue{
			Creator:       oldIssue.Creator,
			Id:            oldIssue.Id,
			Iid:           oldIssue.Iid,
			Title:         oldIssue.Title,
			State:         gitopiatypes.Issue_State(oldIssue.State),
			Description:   oldIssue.Description,
			CommentsCount: oldIssue.CommentsCount,
			RepositoryId:  oldIssue.RepositoryId,
			Labels:        oldIssue.Labels,
			Weight:        oldIssue.Weight,
			Assignees:     oldIssue.Assignees,
			CreatedAt:     oldIssue.CreatedAt,
			UpdatedAt:     oldIssue.UpdatedAt,
			ClosedAt:      oldIssue.ClosedAt,
			ClosedBy:      oldIssue.ClosedBy,
		}

		// Migrate comments
		for _, commentId := range oldIssue.Comments {
			oldComment := commentMap[commentId]

			var attachments []*gitopiatypes.Attachment

			for _, attachmentStr := range oldComment.Attachments {
				var attachment gitopiatypes.Attachment
				if err := json.Unmarshal([]byte(attachmentStr), &attachment); err != nil {
					return gitopiatypes.GenesisState{}, err
				}

				attachments = append(attachments, &attachment)
			}

			commentType := gitopiatypes.CommentTypeNone

			// Set comment type in case of system comment
			if oldComment.System {
				if strings.Contains(oldComment.Body, "assigned to") {
					commentType = gitopiatypes.CommentTypeAddAssignees
				} else if strings.Contains(oldComment.Body, "unassigned") {
					commentType = gitopiatypes.CommentTypeRemoveAssignees
				} else if strings.Contains(oldComment.Body, "added") {
					commentType = gitopiatypes.CommentTypeAddLabels
				} else if strings.Contains(oldComment.Body, "remove") {
					commentType = gitopiatypes.CommentTypeRemoveLabels
				} else if strings.Contains(oldComment.Body, "changed title from") {
					commentType = gitopiatypes.CommentTypeModifiedTitle
				} else if strings.Contains(oldComment.Body, "changed the description") {
					commentType = gitopiatypes.CommentTypeModifiedDescription
				} else if strings.Contains(oldComment.Body, "reopened") {
					commentType = gitopiatypes.CommentTypeIssueOpened
				} else if strings.Contains(oldComment.Body, "closed") {
					commentType = gitopiatypes.CommentTypeIssueClosed
				}
			} else {
				commentType = gitopiatypes.CommentTypeReply
			}

			comment := gitopiatypes.Comment{
				Creator:           oldComment.Creator,
				Id:                oldComment.Id,
				RepositoryId:      issue.RepositoryId,
				ParentIid:         issue.Iid,
				Parent:            gitopiatypes.CommentParentIssue,
				CommentIid:        oldComment.CommentIid,
				Body:              oldComment.Body,
				Attachments:       attachments,
				DiffHunk:          oldComment.DiffHunk,
				Path:              oldComment.Path,
				System:            oldComment.System,
				AuthorAssociation: oldComment.AuthorAssociation,
				CreatedAt:         oldComment.CreatedAt,
				UpdatedAt:         oldComment.UpdatedAt,
				CommentType:       commentType,
			}

			genesisState.CommentList = append(genesisState.CommentList, comment)
		}

		genesisState.IssueList = append(genesisState.IssueList, issue)
	}

	genesisState.IssueCount = state.IssueCount

	for _, oldPullRequest := range state.GetPullRequestList() {
		pullRequest := gitopiatypes.PullRequest{
			Creator:             oldPullRequest.Creator,
			Id:                  oldPullRequest.Id,
			Iid:                 oldPullRequest.Iid,
			Title:               oldPullRequest.Title,
			State:               gitopiatypes.PullRequest_State(oldPullRequest.State),
			Description:         oldPullRequest.Description,
			Locked:              oldPullRequest.Locked,
			CommentsCount:       oldPullRequest.CommentsCount,
			Labels:              oldPullRequest.Labels,
			Assignees:           oldPullRequest.Assignees,
			Reviewers:           oldPullRequest.Reviewers,
			Draft:               oldPullRequest.Draft,
			CreatedAt:           oldPullRequest.CreatedAt,
			UpdatedAt:           oldPullRequest.UpdatedAt,
			ClosedAt:            oldPullRequest.ClosedAt,
			ClosedBy:            oldPullRequest.ClosedBy,
			MergedAt:            oldPullRequest.MergedAt,
			MergedBy:            oldPullRequest.MergedBy,
			MergeCommitSha:      oldPullRequest.MergeCommitSha,
			MaintainerCanModify: oldPullRequest.MaintainerCanModify,
			Head:                (*gitopiatypes.PullRequestHead)(oldPullRequest.Head),
			Base:                (*gitopiatypes.PullRequestBase)(oldPullRequest.Base),
		}

		// Migrate comments
		for _, commentId := range oldPullRequest.Comments {
			oldComment := commentMap[commentId]

			var attachments []*gitopiatypes.Attachment

			for _, attachmentStr := range oldComment.Attachments {
				var attachment gitopiatypes.Attachment
				if err := json.Unmarshal([]byte(attachmentStr), &attachment); err != nil {
					return gitopiatypes.GenesisState{}, err
				}

				attachments = append(attachments, &attachment)
			}

			var commentType gitopiatypes.CommentType

			// Set comment type in case of system comment
			if oldComment.System {
				if strings.Contains(oldComment.Body, "assigned to") {
					commentType = gitopiatypes.CommentTypeAddAssignees
				} else if strings.Contains(oldComment.Body, "unassigned") {
					commentType = gitopiatypes.CommentTypeRemoveAssignees
				} else if strings.Contains(oldComment.Body, "added") {
					commentType = gitopiatypes.CommentTypeAddLabels
				} else if strings.Contains(oldComment.Body, "remove") {
					commentType = gitopiatypes.CommentTypeRemoveLabels
				} else if strings.Contains(oldComment.Body, "changed title from") {
					commentType = gitopiatypes.CommentTypeModifiedTitle
				} else if strings.Contains(oldComment.Body, "changed the description") {
					commentType = gitopiatypes.CommentTypeModifiedDescription
				} else if strings.Contains(oldComment.Body, "reopened") {
					commentType = gitopiatypes.CommentTypePullRequestOpened
				} else if strings.Contains(oldComment.Body, "closed") {
					commentType = gitopiatypes.CommentTypePullRequestClosed
				} else if strings.Contains(oldComment.Body, "merged") {
					commentType = gitopiatypes.CommentTypePullRequestMerged
				} else if strings.Contains(oldComment.Body, "requested review from") {
					commentType = gitopiatypes.CommentTypeAddReviewers
				} else if strings.Contains(oldComment.Body, "removed review request for") {
					commentType = gitopiatypes.CommentTypeRemoveReviewers
				}
			}

			comment := gitopiatypes.Comment{
				Creator:           oldComment.Creator,
				Id:                oldComment.Id,
				RepositoryId:      pullRequest.Base.RepositoryId,
				ParentIid:         pullRequest.Iid,
				Parent:            gitopiatypes.CommentParentPullRequest,
				CommentIid:        oldComment.CommentIid,
				Body:              oldComment.Body,
				Attachments:       attachments,
				DiffHunk:          oldComment.DiffHunk,
				Path:              oldComment.Path,
				System:            oldComment.System,
				AuthorAssociation: oldComment.AuthorAssociation,
				CreatedAt:         oldComment.CreatedAt,
				UpdatedAt:         oldComment.UpdatedAt,
				CommentType:       commentType,
			}

			genesisState.CommentList = append(genesisState.CommentList, comment)
		}

		genesisState.PullRequestList = append(genesisState.PullRequestList, pullRequest)
	}

	genesisState.PullRequestCount = state.PullRequestCount

	return genesisState, nil
}

func GenerateGenesisCmd() *cobra.Command {
	cmd := cobra.Command{
		Use:   "generate-genesis [exported-testnet-state]",
		Short: "Generate Mainnet Genesis file",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := client.GetClientContextFromCmd(cmd)

			blob, err := ioutil.ReadFile(args[0])
			if err != nil {
				return err
			}

			chainID, err := cmd.Flags().GetString(flags.FlagChainID)
			if err != nil {
				return err
			}

			genesisTime, err := cmd.Flags().GetString(flagGenesisTime)
			if err != nil {
				return err
			}

			initialHeight, err := cmd.Flags().GetInt64(flagInitialHeight)
			if err != nil {
				return err
			}

			genesis, err := tmtypes.GenesisDocFromJSON(blob)
			if err != nil {
				return err
			}

			var state genutiltypes.AppMap
			if err := json.Unmarshal(genesis.AppState, &state); err != nil {
				return err
			}

			var (
				testnetGenesis testnettypes.GenesisState
				authGenesis      authtypes.GenesisState
			)

			ctx.Codec.MustUnmarshalJSON(state[gitopiatypes.ModuleName], &testnetGenesis)
			genesisState, err := migrateTestnetState(testnetGenesis)
			if err != nil {
				return err
			}

			ctx.Codec.MustUnmarshalJSON(state[authtypes.ModuleName], &authGenesis)

			var baseAccounts []*codectypes.Any
			for i := range authGenesis.Accounts {
				if authGenesis.Accounts[i].TypeUrl == "/cosmos.auth.v1beta1.BaseAccount" {
					acc := authGenesis.Accounts[i].GetCachedValue().(authtypes.AccountI)
					err = acc.SetSequence(0)
					if err != nil {
						return err
					}
					accAny, err := codectypes.NewAnyWithValue(acc)
					if err != nil {
						return err
					}
					baseAccounts = append(baseAccounts, accAny)
				}
			}
			authGenesis.Accounts = baseAccounts

			var (
				authzGenesis        = authz.DefaultGenesisState()
				bankGenesis         = banktypes.DefaultGenesisState()
				crisisGenesis       = crisistypes.DefaultGenesisState()
				govGenesis          = govv1types.DefaultGenesisState()
				mintGenesis         = minttypes.DefaultGenesisState()
				ibcGenesis          = ibctypes.DefaultGenesisState()
				distributionGenesis = distributiontypes.DefaultGenesisState()
				slashingGenesis     = slashingtypes.DefaultGenesisState()
				genutilGenesis      = genutiltypes.DefaultGenesisState()
				stakingGenesis      = stakingtypes.DefaultGenesisState()
				groupGenesis        = group.NewGenesisState()
				capabilityGenesis   = capabilitytypes.DefaultGenesis()
				evidenceGenesis     = evidencetypes.DefaultGenesisState()
				feegrantGenesis     = feegranttypes.DefaultGenesisState()
				rewardsGenesis      = rewardstypes.DefaultGenesis()
			)

			for _, user := range genesisState.UserList {
				for _, t := range keeper.GitServerTypeUrls {
					a := authz.GrantAuthorization{
						Granter:       user.Creator,
						Grantee:       genesisState.Params.GitServer,
						Authorization: newAnyAuthorization(authz.NewGenericAuthorization(t)),
					}
					authzGenesis.Authorization = append(authzGenesis.Authorization, a)
				}

				for _, t := range keeper.StorageTypeUrls {
					a := authz.GrantAuthorization{
						Granter:       user.Creator,
						Grantee:       genesisState.Params.StorageProvider,
						Authorization: newAnyAuthorization(authz.NewGenericAuthorization(t)),
					}
					authzGenesis.Authorization = append(authzGenesis.Authorization, a)
				}
			}

			for _, dao := range genesisState.DaoList {
				for _, t := range keeper.GitServerTypeUrls {
					a := authz.GrantAuthorization{
						Granter:       dao.Address,
						Grantee:       genesisState.Params.GitServer,
						Authorization: newAnyAuthorization(authz.NewGenericAuthorization(t)),
					}
					authzGenesis.Authorization = append(authzGenesis.Authorization, a)
				}

				for _, t := range keeper.StorageTypeUrls {
					a := authz.GrantAuthorization{
						Granter:       dao.Address,
						Grantee:       genesisState.Params.StorageProvider,
						Authorization: newAnyAuthorization(authz.NewGenericAuthorization(t)),
					}
					authzGenesis.Authorization = append(authzGenesis.Authorization, a)
				}
			}

			fourteenDays := 14 * 24 * time.Hour
			depositParams := govv1types.NewDepositParams(
				sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, math.NewInt(10000000))),
				fourteenDays,
			)
			govGenesis.DepositParams = &depositParams

			votingParams := govv1types.NewVotingParams(fourteenDays)
			govGenesis.VotingParams = &votingParams

			// Disable all transfers
			bankGenesis.Params.DefaultSendEnabled = false
			bankGenesis.DenomMetadata = []banktypes.Metadata{
				{
					Description: "The native staking token of the Gitopia Hub.",
					DenomUnits: []*banktypes.DenomUnit{
						{Denom: params.BaseCoinUnit, Exponent: uint32(0), Aliases: []string{"microlore"}},
						{Denom: params.CoinUnit, Exponent: uint32(6), Aliases: []string{}},
					},
					Base:    params.BaseCoinUnit,
					Display: params.HumanCoinUnit,
					Name:    params.HumanCoinUnit,
					Symbol:  params.HumanCoinUnit,
				},
			}
			bankGenesis.Balances = append(bankGenesis.Balances, banktypes.Balance{
				Address: strategicReserveAddress1,
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(50_000_000_000_000))), // 50M LORE
			}, banktypes.Balance{
				Address: strategicReserveAddress2,
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(50_000_000_000_000))), // 50M LORE
			}, banktypes.Balance{
				Address: strategicReserveAddress3,
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(50_000_000_000_000))), // 50M LORE
			}, banktypes.Balance{
				Address: strategicReserveAddress4,
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(50_000_000_000_000))), // 50M LORE
			}, banktypes.Balance{
				Address: strategicReserveAddress5,
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(50_000_000_000_000))), // 50M LORE
			}, banktypes.Balance{
				Address: feegrantsAddress,
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(1_000_000_000_000))), // 1M LORE of ecosystem incentives for feegrants
			}, banktypes.Balance{
				Address: "gitopia1rtf8ddfa780h4za8j2dss65f7kccurmwktth89",
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(8_571_428_570_000))), // 8,571,428.57 LORE
			}, banktypes.Balance{
				Address: "gitopia1l85lsrnzcfr3llsgs993ceddgqrnutm9ncymrk",
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(1_785_714_290_000))), // 1,785,714.29 LORE
			}, banktypes.Balance{
				Address: "gitopia12zu648n89ve3dg2qul7h28h8fkt6cqp4wt3l4z",
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(14_285_714_290_000))), // 14,285,714.29 LORE
			}, banktypes.Balance{
				Address: "gitopia159a98x95n8uwguhxnf8gnzpy6wj6reu2effl8g",
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(1_785_714_290_000))), // 1,785,714.29 LORE
			}, banktypes.Balance{
				Address: "gitopia1vxh2drxeu5ef4zy8atp59s78shy9dqetn3jedd",
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(892_857_140_000))), // 892,857.14 LORE
			}, banktypes.Balance{
				Address: "gitopia1agd5k6zpksxkw5ufdtf73npluk5nuqa5h5eenr",
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(892_857_140_000))), // 892,857.14 LORE
			}, banktypes.Balance{
				Address: "gitopia1za95wp6a7qhdyu099797dtfsxfmgs0tzwata9p",
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(7_142_857_140_000))), // 7,142,857.14 LORE
			}, banktypes.Balance{
				Address: earlySupportersMultiSigAddress,
				// 18,571,428.57 LORE
				// 17,857,142.86 LORE + 714,285.71 LORE
				Coins: sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(25_714_285_710_000))),
			}, banktypes.Balance{
				Address: strategicPartnersMultiSigAddress,
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(STRATEGIC_PARTNERS_AMOUNT))),
			}, banktypes.Balance{
				Address: advisorsMultiSigAddress,
				Coins:   sdk.NewCoins(sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(ADVISORS_AMOUNT))),
			})

			crisisGenesis.ConstantFee = sdk.NewCoin(params.BaseCoinUnit, sdk.NewInt(1000))

			mintGenesis.Minter.Inflation = sdk.NewDecWithPrec(35, 2)
			mintGenesis.Params = minttypes.Params{
				MintDenom:           params.BaseCoinUnit,
				InflationRateChange: sdk.NewDecWithPrec(20, 2),
				InflationMax:        sdk.NewDecWithPrec(45, 2),
				InflationMin:        sdk.NewDecWithPrec(25, 2),
				GoalBonded:          sdk.NewDecWithPrec(67, 2),
				BlocksPerYear:       uint64(60 * 60 * 8766 / 1.5), // assuming 1.5 second block times
			}

			distributionGenesis.Params = distributiontypes.Params{
				CommunityTax:        sdk.NewDecWithPrec(5, 2), // 5%
				BaseProposerReward:  sdk.NewDecWithPrec(1, 2), // 1%
				BonusProposerReward: sdk.NewDecWithPrec(4, 2), // 4%
				WithdrawAddrEnabled: true,
			}

			stakingGenesis.Params = stakingtypes.NewParams(
				stakingtypes.DefaultUnbondingTime,
				stakingtypes.DefaultMaxValidators,
				stakingtypes.DefaultMaxEntries,
				stakingtypes.DefaultHistoricalEntries,
				params.BaseCoinUnit,
				stakingtypes.DefaultMinCommissionRate,
			)

			rewardsGenesis.Params = rewardstypes.NewParams(rewardsServiceAddress, &rewardstypes.RewardSeries{
				SeriesOne: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(2_000_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesTwo: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(1_700_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesThree: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(1_400_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesFour: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(1_100_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesFive: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(800_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesSix: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(500_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
				SeriesSeven: &rewardstypes.RewardPool{
					TotalAmount:   sdk.NewCoin(params.BaseCoinUnit, math.NewInt(250_000_000_000)),
					ClaimedAmount: sdk.NewCoin(params.BaseCoinUnit, math.NewInt(0)),
				},
			})

			state[authtypes.ModuleName] = ctx.Codec.MustMarshalJSON(&authGenesis)
			state[authz.ModuleName] = ctx.Codec.MustMarshalJSON(authzGenesis)
			state[banktypes.ModuleName] = ctx.Codec.MustMarshalJSON(bankGenesis)
			state[crisistypes.ModuleName] = ctx.Codec.MustMarshalJSON(crisisGenesis)
			state[govtypes.ModuleName] = ctx.Codec.MustMarshalJSON(govGenesis)
			state[distributiontypes.ModuleName] = ctx.Codec.MustMarshalJSON(distributionGenesis)
			state[minttypes.ModuleName] = ctx.Codec.MustMarshalJSON(mintGenesis)
			state[slashingtypes.ModuleName] = ctx.Codec.MustMarshalJSON(slashingGenesis)
			state[genutiltypes.ModuleName] = ctx.Codec.MustMarshalJSON(genutilGenesis)
			state[stakingtypes.ModuleName] = ctx.Codec.MustMarshalJSON(stakingGenesis)
			state[ibchost.ModuleName] = ctx.Codec.MustMarshalJSON(ibcGenesis)
			state[gitopiatypes.ModuleName] = ctx.Codec.MustMarshalJSON(&genesisState)
			state[group.ModuleName] = ctx.Codec.MustMarshalJSON(groupGenesis)
			state[capabilitytypes.ModuleName] = ctx.Codec.MustMarshalJSON(capabilityGenesis)
			state[evidencetypes.ModuleName] = ctx.Codec.MustMarshalJSON(evidenceGenesis)
			state[feegranttypes.ModuleName] = ctx.Codec.MustMarshalJSON(feegrantGenesis)
			state[rewardstypes.ModuleName] = ctx.Codec.MustMarshalJSON(rewardsGenesis)

			genesis.AppState, err = json.Marshal(state)
			if err != nil {
				return err
			}

			if genesisTime != "" {
				var t time.Time
				if err := t.UnmarshalText([]byte(genesisTime)); err != nil {
					return err
				}

				genesis.GenesisTime = t
			}
			if chainID != "" {
				genesis.ChainID = chainID
			}

			genesis.InitialHeight = initialHeight

			genesis.Validators = []tmtypes.GenesisValidator{}

			blob, err = tmjson.Marshal(genesis)
			if err != nil {
				return err
			}

			sortedBlob, err := sdk.SortJSON(blob)
			if err != nil {
				return err
			}

			fmt.Println(string(sortedBlob))
			return nil
		},
	}

	cmd.Flags().String(flags.FlagChainID, "gitopia", "set chain id")
	cmd.Flags().String(flagGenesisTime, time.Now().UTC().Format(time.RFC3339Nano), "set genesis time")
	cmd.Flags().Int64(flagInitialHeight, 1, "set the initial height")

	return &cmd
}
