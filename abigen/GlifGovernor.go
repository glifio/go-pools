// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// GlifGovernorMetaData contains all meta data concerning the GlifGovernor contract.
var GlifGovernorMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVotes\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_initialQuorumNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"_initialVotingDelay\",\"type\":\"uint48\"},{\"internalType\":\"uint32\",\"name\":\"_initialVotingPeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"_initialProposalThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint48\",\"name\":\"_initialVoteExtension\",\"type\":\"uint48\"},{\"internalType\":\"contractTimelockController\",\"name\":\"_timelock\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"CheckpointUnorderedInsertion\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorAlreadyCastVote\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorAlreadyQueuedProposal\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorCountingFractional_NoVoteWeight\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorCountingFractional__InvalidVoteData\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorCountingFractional__VoteWeightExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorDisabledDeposit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"votes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"threshold\",\"type\":\"uint256\"}],\"name\":\"GovernorInsufficientProposerVotes\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"targets\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"calldatas\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"values\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidProposalLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"quorumNumerator\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quorumDenominator\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidQuorumFraction\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"}],\"name\":\"GovernorInvalidSignature\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorInvalidVoteType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"votingPeriod\",\"type\":\"uint256\"}],\"name\":\"GovernorInvalidVotingPeriod\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNonexistentProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"GovernorNotQueuedProposal\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyExecutor\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"GovernorOnlyProposer\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GovernorQueueNotImplemented\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"}],\"name\":\"GovernorRestrictedProposer\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"current\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"expectedStates\",\"type\":\"bytes32\"}],\"name\":\"GovernorUnexpectedProposalState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"currentNonce\",\"type\":\"uint256\"}],\"name\":\"InvalidAccountNonce\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidShortString\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueEmpty\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueueFull\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint8\",\"name\":\"bits\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintDowncast\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"str\",\"type\":\"string\"}],\"name\":\"StringTooLong\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"EIP712DomainChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"oldVoteExtension\",\"type\":\"uint64\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"newVoteExtension\",\"type\":\"uint64\"}],\"name\":\"LateQuorumVoteExtensionSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalCanceled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"proposer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"string[]\",\"name\":\"signatures\",\"type\":\"string[]\"},{\"indexed\":false,\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteStart\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"voteEnd\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"ProposalCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint64\",\"name\":\"extendedDeadline\",\"type\":\"uint64\"}],\"name\":\"ProposalExtended\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"etaSeconds\",\"type\":\"uint256\"}],\"name\":\"ProposalQueued\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldProposalThreshold\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"ProposalThresholdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldQuorumNumerator\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"QuorumNumeratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldTimelock\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"TimelockChange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"VoteCast\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"weight\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"VoteCastWithParams\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingDelay\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingDelay\",\"type\":\"uint256\"}],\"name\":\"VotingDelaySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldVotingPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newVotingPeriod\",\"type\":\"uint256\"}],\"name\":\"VotingPeriodSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CLOCK_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"COUNTING_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXTENDED_BALLOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"cancel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"}],\"name\":\"castVote\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"}],\"name\":\"castVoteWithReason\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"support\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"voter\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"reason\",\"type\":\"string\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"castVoteWithReasonAndParamsBySig\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"clock\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"eip712Domain\",\"outputs\":[{\"internalType\":\"bytes1\",\"name\":\"fields\",\"type\":\"bytes1\"},{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"version\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"verifyingContract\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"extensions\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"execute\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"getVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"name\":\"getVotesWithParams\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasVoted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"hashProposal\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lateQuorumVoteExtension\",\"outputs\":[{\"internalType\":\"uint48\",\"name\":\"\",\"type\":\"uint48\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155BatchReceived\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC1155Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"onERC721Received\",\"outputs\":[{\"internalType\":\"bytes4\",\"name\":\"\",\"type\":\"bytes4\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalDeadline\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalEta\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalNeedsQueuing\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalProposer\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposalThreshold\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"proposalVotes\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"againstVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"forVotes\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"abstainVotes\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"}],\"name\":\"propose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"targets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes[]\",\"name\":\"calldatas\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"descriptionHash\",\"type\":\"bytes32\"}],\"name\":\"queue\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumDenominator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"timepoint\",\"type\":\"uint256\"}],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"quorumNumerator\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"relay\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newVoteExtension\",\"type\":\"uint48\"}],\"name\":\"setLateQuorumVoteExtension\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newProposalThreshold\",\"type\":\"uint256\"}],\"name\":\"setProposalThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint48\",\"name\":\"newVotingDelay\",\"type\":\"uint48\"}],\"name\":\"setVotingDelay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newVotingPeriod\",\"type\":\"uint32\"}],\"name\":\"setVotingPeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"}],\"name\":\"state\",\"outputs\":[{\"internalType\":\"enumIGovernor.ProposalState\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"timelock\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC5805\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newQuorumNumerator\",\"type\":\"uint256\"}],\"name\":\"updateQuorumNumerator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractTimelockController\",\"name\":\"newTimelock\",\"type\":\"address\"}],\"name\":\"updateTimelock\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"proposalId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"voteWeightCast\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingDelay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"votingPeriod\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x61018060405234801562000011575f80fd5b5060405162005ef338038062005ef3833981016040819052620000349162000848565b8185858589858c6040518060400160405280601081526020016f476c696620476f7665726e6f7220763160801b8152508062000075620001a060201b60201c565b62000081825f620001bb565b6101205262000092816001620001bb565b61014052815160208084019190912060e052815190820120610100524660a0526200011f60e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c052600362000136828262000977565b50506001600160a01b0316610160526200015081620001f3565b506200015c816200025c565b506200016883620002ff565b620001738262000365565b6200017e8162000406565b50505062000192816200044760201b60201c565b505050505050505062000aec565b6040805180820190915260018152603160f81b602082015290565b5f602083511015620001da57620001d283620004cc565b9050620001ed565b81620001e7848262000977565b5060ff90505b92915050565b600954604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a1600980546001600160a01b0319166001600160a01b0392909216919091179055565b6064808211156200028f5760405163243e544560e01b815260048101839052602481018290526044015b60405180910390fd5b5f6200029a6200050e565b9050620002bf620002aa62000529565b620002b585620005aa565b600b9190620005e3565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b600d546040805165ffffffffffff928316815291831660208301527fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93910160405180910390a1600d805465ffffffffffff191665ffffffffffff92909216919091179055565b8063ffffffff165f036200038f5760405163f1cfbf0560e01b81525f600482015260240162000286565b600d546040805163ffffffff66010000000000009093048316815291831660208301527f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828910160405180910390a1600d805463ffffffff90921666010000000000000263ffffffff60301b19909216919091179055565b600c5460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461910160405180910390a1600c55565b600d546040805165ffffffffffff6a01000000000000000000009093048316815291831660208301527f7ca4ac117ed3cdce75c1161d8207c440389b1a15d69d096831664657c07dafc2910160405180910390a1600d805465ffffffffffff9092166a01000000000000000000000265ffffffffffff60501b19909216919091179055565b5f80829050601f81511115620004f9578260405163305a27a960e01b815260040162000286919062000a3f565b8051620005068262000a8c565b179392505050565b5f6200051b600b620005ff565b6001600160d01b0316905090565b5f620005356101605190565b6001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801562000591575060408051601f3d908101601f191682019092526200058e9181019062000ab0565b60015b620005a557620005a06200064c565b905090565b919050565b5f6001600160d01b03821115620005df576040516306dfcc6560e41b815260d060048201526024810183905260440162000286565b5090565b5f80620005f285858562000658565b915091505b935093915050565b80545f908015620006435762000629836200061c60018462000acc565b5f91825260209091200190565b54660100000000000090046001600160d01b031662000645565b5f5b9392505050565b5f620005a043620007e6565b82545f908190801562000788575f62000678876200061c60018562000acc565b60408051808201909152905465ffffffffffff80821680845266010000000000009092046001600160d01b031660208401529192509087161015620006d057604051632520601d60e01b815260040160405180910390fd5b805165ffffffffffff808816911603620007245784620006f7886200061c60018662000acc565b80546001600160d01b039290921666010000000000000265ffffffffffff90921691909117905562000777565b6040805180820190915265ffffffffffff80881682526001600160d01b0380881660208085019182528b54600181018d555f8d815291909120945191519092166601000000000000029216919091179101555b602001519250839150620005f79050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a555f8a8152918220955192519093166601000000000000029190931617920191909155905081620005f7565b5f65ffffffffffff821115620005df576040516306dfcc6560e41b8152603060048201526024810183905260440162000286565b6001600160a01b03811681146200082f575f80fd5b50565b805165ffffffffffff81168114620005a5575f80fd5b5f805f805f805f60e0888a0312156200085f575f80fd5b87516200086c816200081a565b60208901519097509550620008846040890162000832565b9450606088015163ffffffff811681146200089d575f80fd5b60808901519094509250620008b560a0890162000832565b915060c0880151620008c7816200081a565b8091505092959891949750929550565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200090057607f821691505b6020821081036200091f57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f82111562000972575f81815260208120601f850160051c810160208610156200094d5750805b601f850160051c820191505b818110156200096e5782815560010162000959565b5050505b505050565b81516001600160401b03811115620009935762000993620008d7565b620009ab81620009a48454620008eb565b8462000925565b602080601f831160018114620009e1575f8415620009c95750858301515b5f19600386901b1c1916600185901b1785556200096e565b5f85815260208120601f198616915b8281101562000a1157888601518255948401946001909101908401620009f0565b508582101562000a2f57878501515f19600388901b60f8161c191681555b5050505050600190811b01905550565b5f6020808352835180828501525f5b8181101562000a6c5785810183015185820160400152820162000a4e565b505f604082860101526040601f19601f8301168501019250505092915050565b805160208083015191908110156200091f575f1960209190910360031b1b16919050565b5f6020828403121562000ac1575f80fd5b620006458262000832565b81810381811115620001ed57634e487b7160e01b5f52601160045260245ffd5b60805160a05160c05160e0516101005161012051610140516101605161538f62000b645f395f8181610ab401528181610da5015281816110bc0152818161234b0152613b5f01525f611ec401525f611e9801525f6136a501525f61367d01525f6135d801525f61360201525f61362c015261538f5ff3fe608060405260043610610327575f3560e01c80637b3c71d3116101a7578063bc197c81116100e7578063deaaa7cc11610092578063ece40cc11161006d578063ece40cc114610a49578063f23a6e6114610a68578063f8ce560a14610a87578063fc0c546a14610aa6575f80fd5b8063deaaa7cc146109d8578063e540d01d14610a0b578063eb9019d414610a2a575f80fd5b8063c59057e4116100c2578063c59057e414610988578063d33219b4146109a7578063dd4e2ba5146109c4575f80fd5b8063bc197c8114610937578063c01f9e3714610956578063c28bc2fa14610975575f80fd5b806397c3d33411610152578063a890c9101161012d578063a890c910146108af578063a9a95294146108ce578063ab58fb8e146108ed578063b58131b014610923575f80fd5b806397c3d334146108695780639a802a6d1461087c578063a7713a701461089b575f80fd5b806384b0196e1161018257806384b0196e1461080f5780638ff262e31461083657806391ddadf414610855575f80fd5b80637b3c71d3146107b25780637d5e81e2146107d15780637ecebe00146107f0575f80fd5b8063342cfab611610272578063544ffc9c1161021d5780635b8d0e0d116101f85780635b8d0e0d146107365780635f398a141461075557806360c4247f146107745780637905188714610793575f80fd5b8063544ffc9c1461066c57806354fd4d50146106d25780635678138814610717575f80fd5b8063438596321161024d578063438596321461061a578063452115d6146106395780634bf5d7e914610658575f80fd5b8063342cfab61461059b5780633932abb1146105d25780633e4f49e6146105ee575f80fd5b8063160cbed7116102d25780632fe3e261116102ad5780632fe3e2611461050c57806332b8113e1461053f578063330df7ff1461057c575f80fd5b8063160cbed7146104a05780632656227d146104bf5780632d63f693146104d2575f80fd5b806306fdde031161030257806306fdde03146103e3578063143489d014610404578063150b7a0214610450575f80fd5b806301ffc9a71461036257806302a251a31461039657806306f3f9e6146103c4575f80fd5b3661035e5730610335610ad8565b6001600160a01b03161461035c57604051637485328f60e11b815260040160405180910390fd5b005b5f80fd5b34801561036d575f80fd5b5061038161037c3660046142df565b610af0565b60405190151581526020015b60405180910390f35b3480156103a1575f80fd5b50600d546601000000000000900463ffffffff165b60405190815260200161038d565b3480156103cf575f80fd5b5061035c6103de36600461431e565b610bd4565b3480156103ee575f80fd5b506103f7610be8565b60405161038d9190614382565b34801561040f575f80fd5b5061043861041e36600461431e565b5f908152600460205260409020546001600160a01b031690565b6040516001600160a01b03909116815260200161038d565b34801561045b575f80fd5b5061046f61046a36600461446d565b610c78565b6040517fffffffff00000000000000000000000000000000000000000000000000000000909116815260200161038d565b3480156104ab575f80fd5b506103b66104ba366004614638565b610cb0565b6103b66104cd366004614638565b610cc5565b3480156104dd575f80fd5b506103b66104ec36600461431e565b5f90815260046020526040902054600160a01b900465ffffffffffff1690565b348015610517575f80fd5b506103b67f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a81181565b34801561054a575f80fd5b50600d546a0100000000000000000000900465ffffffffffff165b60405165ffffffffffff909116815260200161038d565b348015610587575f80fd5b5061035c6105963660046146d5565b610cda565b3480156105a6575f80fd5b506105ba6105b53660046146f0565b610ceb565b6040516001600160801b03909116815260200161038d565b3480156105dd575f80fd5b50600d5465ffffffffffff166103b6565b3480156105f9575f80fd5b5061060d61060836600461431e565b610d39565b60405161038d9190614752565b348015610625575f80fd5b506103816106343660046146f0565b610d43565b348015610644575f80fd5b506103b6610653366004614638565b610d8c565b348015610663575f80fd5b506103f7610da1565b348015610677575f80fd5b506106b761068636600461431e565b5f90815260076020526040902080546001909101546001600160801b0380831693600160801b909304811692911690565b6040805193845260208401929092529082015260600161038d565b3480156106dd575f80fd5b5060408051808201909152600181527f310000000000000000000000000000000000000000000000000000000000000060208201526103f7565b348015610722575f80fd5b506103b6610731366004614770565b610e61565b348015610741575f80fd5b506103b66107503660046147df565b610e88565b348015610760575f80fd5b506103b661076f366004614892565b610eb5565b34801561077f575f80fd5b506103b661078e36600461431e565b610f08565b34801561079e575f80fd5b5061035c6107ad3660046146d5565b610f97565b3480156107bd575f80fd5b506103b66107cc366004614910565b610fa8565b3480156107dc575f80fd5b506103b66107eb366004614966565b610ff8565b3480156107fb575f80fd5b506103b661080a366004614a13565b61100d565b34801561081a575f80fd5b5061082361103d565b60405161038d9796959493929190614a67565b348015610841575f80fd5b506103b6610850366004614af0565b61109b565b348015610860575f80fd5b506105656110b9565b348015610874575f80fd5b5060646103b6565b348015610887575f80fd5b506103b6610896366004614b3e565b611140565b3480156108a6575f80fd5b506103b661115d565b3480156108ba575f80fd5b5061035c6108c9366004614a13565b611176565b3480156108d9575f80fd5b506103816108e836600461431e565b611187565b3480156108f8575f80fd5b506103b661090736600461431e565b5f9081526004602052604090206001015465ffffffffffff1690565b34801561092e575f80fd5b506103b661118f565b348015610942575f80fd5b5061046f610951366004614b93565b611199565b348015610961575f80fd5b506103b661097036600461431e565b6111c9565b61035c610983366004614c20565b6111d3565b348015610993575f80fd5b506103b66109a2366004614638565b611260565b3480156109b2575f80fd5b506009546001600160a01b0316610438565b3480156109cf575f80fd5b506103f761126d565b3480156109e3575f80fd5b506103b67ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d781565b348015610a16575f80fd5b5061035c610a25366004614c60565b61128d565b348015610a35575f80fd5b506103b6610a44366004614c83565b61129e565b348015610a54575f80fd5b5061035c610a6336600461431e565b6112ba565b348015610a73575f80fd5b5061046f610a82366004614cad565b6112cb565b348015610a92575f80fd5b506103b6610aa136600461431e565b6112fb565b348015610ab1575f80fd5b507f0000000000000000000000000000000000000000000000000000000000000000610438565b5f610aeb6009546001600160a01b031690565b905090565b5f7fffffffff0000000000000000000000000000000000000000000000000000000082167f65455a86000000000000000000000000000000000000000000000000000000001480610b8257507fffffffff0000000000000000000000000000000000000000000000000000000082167f4e2312e000000000000000000000000000000000000000000000000000000000145b80610bce57507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff000000000000000000000000000000000000000000000000000000008316145b92915050565b610bdc611305565b610be58161139c565b50565b606060038054610bf790614d11565b80601f0160208091040260200160405190810160405280929190818152602001828054610c2390614d11565b8015610c6e5780601f10610c4557610100808354040283529160200191610c6e565b820191905f5260205f20905b815481529060010190602001808311610c5157829003601f168201915b5050505050905090565b5f610ca7610c8e866001600160a01b031661144a565b610ca0866001600160a01b031661144a565b8585611485565b95945050505050565b5f610ca7610cbd866114e0565b8585856115a2565b5f610ca7610cd2866114e0565b858585611687565b610ce2611305565b610be58161182d565b5f610d3283610d02846001600160a01b031661144a565b5f9182526008602090815260408084206001600160a01b039390931684529190529020546001600160801b031690565b9392505050565b5f610bce826118c8565b5f610d3283610d5a846001600160a01b031661144a565b5f9182526008602090815260408084206001600160a01b039390931684529190529020546001600160801b0316151590565b5f610ca7610d99866114e0565b858585611a33565b60607f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634bf5d7e96040518163ffffffff1660e01b81526004015f60405180830381865afa925050508015610e2057506040513d5f823e601f3d908101601f19168201604052610e1d9190810190614d49565b60015b610e5c575060408051808201909152601d81527f6d6f64653d626c6f636b6e756d6265722666726f6d3d64656661756c74000000602082015290565b919050565b5f80339050610e8084828560405180602001604052805f815250611aae565b949350505050565b5f610ea98888610ea0896001600160a01b031661144a565b88888888611acf565b98975050505050505050565b5f80339050610efd87828888888080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508a9250611c44915050565b979650505050505050565b600b80545f918290610f1b600184614dc6565b81548110610f2b57610f2b614dd9565b5f918252602090912001805490915065ffffffffffff811690660100000000000090046001600160d01b0316858211610f70576001600160d01b031695945050505050565b610f84610f7c87611c63565b600b90611c99565b6001600160d01b03169695505050505050565b610f9f611305565b610be581611d4b565b5f80339050610fee86828787878080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250611aae92505050565b9695505050505050565b5f610ca7611005866114e0565b858585611db1565b5f610bce611023836001600160a01b031661144a565b6001600160a01b03165f9081526002602052604090205490565b5f6060805f805f606061104e611e91565b611056611ebd565b604080515f808252602082019092527f0f000000000000000000000000000000000000000000000000000000000000009b939a50919850469750309650945092509050565b5f610ca785856110b3866001600160a01b031661144a565b85611eea565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166391ddadf46040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611134575060408051601f3d908101601f1916820190925261113191810190614ded565b60015b610e5c57610aeb611fd2565b5f610e80611156856001600160a01b031661144a565b8484611fdc565b5f611168600b611fe8565b6001600160d01b0316905090565b61117e611305565b610be581612022565b5f6001610bce565b5f610aeb600c5490565b5f610fee6111af876001600160a01b031661144a565b6111c1876001600160a01b031661144a565b8686866120a3565b5f610bce826120ff565b6111db611305565b5f806111ef866001600160a01b031661144a565b6001600160a01b0316858585604051611209929190614e08565b5f6040518083038185875af1925050503d805f8114611243576040519150601f19603f3d011682016040523d82523d5f602084013e611248565b606091505b50915091506112578282612127565b50505050505050565b5f610ca785858585612143565b606060405180606001604052806032815260200161532860329139905090565b611295611305565b610be58161217c565b5f610d326112b4846001600160a01b031661144a565b8361224c565b6112c2611305565b610be58161226b565b5f610fee6112e1876001600160a01b031661144a565b6112f3876001600160a01b031661144a565b8686866122ac565b5f610bce82612308565b3361130e610ad8565b6001600160a01b031614611355576040517f47096e470000000000000000000000000000000000000000000000000000000081523360048201526024015b60405180910390fd5b3061135e610ad8565b6001600160a01b03161461139a575f803660405161137d929190614e08565b604051809103902090505b8061139360056123c8565b0361138857505b565b6064808211156113e2576040517f243e5445000000000000000000000000000000000000000000000000000000008152600481018390526024810182905260440161134c565b5f6113eb61115d565b905061140a6113f86110b9565b6114018561245d565b600b9190612490565b505060408051828152602081018590527f0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997910160405180910390a1505050565b5f805f611456846124aa565b915091508161146757509192915050565b5f80611472836124f6565b9150915081610ca7575093949350505050565b5f3061148f610ad8565b6001600160a01b0316146114b657604051637485328f60e11b815260040160405180910390fd5b507f150b7a0200000000000000000000000000000000000000000000000000000000949350505050565b60605f825167ffffffffffffffff8111156114fd576114fd6143a8565b604051908082528060200260200182016040528015611526578160200160208202803683370190505b5090505f5b835181101561159b5761155f84828151811061154957611549614dd9565b60200260200101516001600160a01b031661144a565b82828151811061157157611571614dd9565b6001600160a01b03909216602092830291909101909101528061159381614e17565b91505061152b565b5092915050565b5f806115b086868686611260565b90506115c5816115c06004612566565b612588565b505f6115d482888888886125de565b905065ffffffffffff81161561164b575f82815260046020908152604091829020600101805465ffffffffffff191665ffffffffffff85169081179091558251858152918201527f9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892910160405180910390a161167d565b6040517f90884a4600000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5095945050505050565b5f8061169586868686611260565b90506116b5816116a56005612566565b6116af6004612566565b17612588565b505f81815260046020526040902080547fff00ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e0100000000000000000000000000000000000000000000000000000000000017905530611713610ad8565b6001600160a01b0316146117a4575f5b86518110156117a257306001600160a01b031687828151811061174857611748614dd9565b60200260200101516001600160a01b0316036117925761179285828151811061177357611773614dd9565b60200260200101518051906020012060056125f490919063ffffffff16565b61179b81614e17565b9050611723565b505b6117b1818787878761267d565b306117ba610ad8565b6001600160a01b0316141580156117e657506005546001600160801b03808216600160801b9092041614155b156117f0575f6005555b6040518181527f712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f906020015b60405180910390a195945050505050565b600d546040805165ffffffffffff6a01000000000000000000009093048316815291831660208301527f7ca4ac117ed3cdce75c1161d8207c440389b1a15d69d096831664657c07dafc2910160405180910390a1600d805465ffffffffffff9092166a0100000000000000000000027fffffffffffffffffffffffffffffffff000000000000ffffffffffffffffffff909216919091179055565b5f806118d383612699565b905060058160078111156118e9576118e961471e565b146118f45792915050565b5f838152600a6020526040908190205460095491517f584b153e0000000000000000000000000000000000000000000000000000000081526004810182905290916001600160a01b03169063584b153e90602401602060405180830381865afa158015611963573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906119879190614e2f565b15611996575060059392505050565b6009546040517f2ab0f529000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b0390911690632ab0f52990602401602060405180830381865afa1580156119f6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611a1a9190614e2f565b15611a29575060079392505050565b5060029392505050565b5f80611a4186868686611260565b9050611a50816115c05f612566565b505f818152600460205260409020546001600160a01b03163314611aa2576040517f233d98e300000000000000000000000000000000000000000000000000000000815233600482015260240161134c565b610fee8686868661283a565b5f610ca785858585611aca60408051602081019091525f815290565b611c44565b5f80611bb087611baa7f3e83946653575f9a39005e1545185629e92736b7528ab20ca3816f315424a8118c8c8c611b228e6001600160a01b03165f90815260026020526040902080546001810190915590565b8d8d604051611b32929190614e08565b60405180910390208c80519060200120604051602001611b8f9796959493929190968752602087019590955260ff9390931660408601526001600160a01b03919091166060850152608084015260a083015260c082015260e00190565b6040516020818303038152906040528051906020012061284f565b85612896565b905080611bf4576040517f94ab6c070000000000000000000000000000000000000000000000000000000081526001600160a01b038816600482015260240161134c565b611c3789888a89898080601f0160208091040260200160405190810160405280939291908181526020018383808284375f920191909152508b9250611c44915050565b9998505050505050505050565b5f610fee86611c5b876001600160a01b031661144a565b8686866128eb565b5f65ffffffffffff821115611c95576040516306dfcc6560e41b8152603060048201526024810183905260440161134c565b5090565b81545f9081816005811115611cf5575f611cb2846129e1565b611cbc9085614dc6565b5f8881526020902090915081015465ffffffffffff9081169087161015611ce557809150611cf3565b611cf0816001614e4e565b92505b505b5f611d0287878585612ac5565b90508015611d3f57611d2687611d19600184614dc6565b5f91825260209091200190565b54660100000000000090046001600160d01b0316610efd565b5f979650505050505050565b600d546040805165ffffffffffff928316815291831660208301527fc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93910160405180910390a1600d805465ffffffffffff191665ffffffffffff92909216919091179055565b5f33611dbd8184612b24565b611dfe576040517fd9b395570000000000000000000000000000000000000000000000000000000081526001600160a01b038216600482015260240161134c565b5f611e24826001611e0d6110b9565b611e179190614e61565b65ffffffffffff1661129e565b90505f611e2f61118f565b905080821015611e84576040517fc242ee160000000000000000000000000000000000000000000000000000000081526001600160a01b0384166004820152602481018390526044810182905260640161134c565b610ea98888888887612c71565b6060610aeb7f00000000000000000000000000000000000000000000000000000000000000005f612f27565b6060610aeb7f00000000000000000000000000000000000000000000000000000000000000006001612f27565b5f80611f7484611baa7ff2aad550cf55f045cb27e9c559f9889fdfb6e6cdaa032301d6ea397784ae51d7898989611f3d8b6001600160a01b03165f90815260026020526040902080546001810190915590565b60408051602081019690965285019390935260ff90911660608401526001600160a01b0316608083015260a082015260c001611b8f565b905080611fb8576040517f94ab6c070000000000000000000000000000000000000000000000000000000081526001600160a01b038516600482015260240161134c565b610fee86858760405180602001604052805f815250611aae565b5f610aeb43611c63565b5f610e80848484612fd0565b80545f90801561201a5761200183611d19600184614dc6565b54660100000000000090046001600160d01b0316610d32565b5f9392505050565b600954604080516001600160a01b03928316815291831660208301527f08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401910160405180910390a1600980547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b0392909216919091179055565b5f306120ad610ad8565b6001600160a01b0316146120d457604051637485328f60e11b815260040160405180910390fd5b507fbc197c810000000000000000000000000000000000000000000000000000000095945050505050565b5f610bce61210c83612fed565b5f848152600e602052604090205465ffffffffffff16613046565b60608261213c576121378261305b565b610bce565b5080610bce565b5f8484848460405160200161215b9493929190614f07565b60408051601f19818403018152919052805160209091012095945050505050565b8063ffffffff165f036121bd576040517ff1cfbf050000000000000000000000000000000000000000000000000000000081525f600482015260240161134c565b600d546040805163ffffffff66010000000000009093048316815291831660208301527f7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828910160405180910390a1600d805463ffffffff9092166601000000000000027fffffffffffffffffffffffffffffffffffffffffffff00000000ffffffffffff909216919091179055565b5f610d32838361226660408051602081019091525f815290565b612fd0565b600c5460408051918252602082018390527fccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461910160405180910390a1600c55565b5f306122b6610ad8565b6001600160a01b0316146122dd57604051637485328f60e11b815260040160405180910390fd5b507ff23a6e610000000000000000000000000000000000000000000000000000000095945050505050565b5f606461231483610f08565b6040517f8e539e8c000000000000000000000000000000000000000000000000000000008152600481018590526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001690638e539e8c90602401602060405180830381865afa158015612390573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906123b49190614f51565b6123be9190614f68565b610bce9190614f93565b80545f906001600160801b0380821691600160801b9004168103612418576040517f75e52f4f00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160801b038181165f908152600185810160205260408220805492905585546fffffffffffffffffffffffffffffffff19169301909116919091179092555090565b5f6001600160d01b03821115611c95576040516306dfcc6560e41b815260d060048201526024810183905260440161134c565b5f8061249d85858561309d565b915091505b935093915050565b5f8073ffffffffffffffffffffffff0000000000000000831673ff0000000000000000000000000000000000000081036124f0576001925067ffffffffffffffff841691505b50915091565b5f80825f526016600a60205f73fe000000000000000000000000000000000000025afa91505f51806001600160a01b031691508060a01c61ffff16905061040a8114612543575f92505f91505b5081158061255257503d601614155b1561256157505f928392509050565b915091565b5f8160078111156125795761257961471e565b600160ff919091161b92915050565b5f8061259384610d39565b90505f836125a083612566565b1603610d32578381846040517f31b75e4d00000000000000000000000000000000000000000000000000000000815260040161134c93929190614fb2565b5f610fee866125ec876114e0565b868686613238565b81546001600160801b03600160801b820481169181166001830190911603612648576040517f8acb5f2700000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6001600160801b038082165f90815260018086016020526040909120939093558354919092018216600160801b029116179055565b6126928561268a866114e0565b85858561340e565b5050505050565b5f818152600460205260408120805460ff7e0100000000000000000000000000000000000000000000000000000000000082048116917f0100000000000000000000000000000000000000000000000000000000000000900416811561270457506007949350505050565b801561271557506002949350505050565b5f85815260046020526040812054600160a01b900465ffffffffffff169050805f03612770576040517f6ad060750000000000000000000000000000000000000000000000000000000081526004810187905260240161134c565b5f6127796110b9565b65ffffffffffff16905080821061279657505f9695505050505050565b5f6127a0886111c9565b90508181106127b757506001979650505050505050565b6127c0886134b1565b15806127ec57505f888152600760205260409020546001600160801b03808216600160801b9092041611155b156127ff57506003979650505050505050565b5f8881526004602052604090206001015465ffffffffffff165f0361282c57506004979650505050505050565b506005979650505050505050565b5f610ca7612847866114e0565b85858561351c565b5f610bce61285b6135cc565b836040517f19010000000000000000000000000000000000000000000000000000000000008152600281019290925260228201526042902090565b5f805f6128a385856136f5565b5090925090505f8160038111156128bc576128bc61471e565b1480156128da5750856001600160a01b0316826001600160a01b0316145b80610fee5750610fee86868661373e565b5f806128fa878787878761385b565b5f888152600e602052604090205490915065ffffffffffff161580156129245750612924876134b1565b15610fee57600d545f906a0100000000000000000000900465ffffffffffff1661294c6110b9565b6129569190614fd4565b9050612961886111c9565b8165ffffffffffff1611156129ae5760405165ffffffffffff8216815288907f541f725fb9f7c98a30cc9c0ff32fbb14358cd7159c847a3aa20a2bdc442ba5119060200160405180910390a25b5f888152600e60205260409020805465ffffffffffff191665ffffffffffff929092169190911790559695505050505050565b5f815f036129f057505f919050565b5f60016129fc8461394a565b901c6001901b90506001818481612a1557612a15614f7f565b048201901c90506001818481612a2d57612a2d614f7f565b048201901c90506001818481612a4557612a45614f7f565b048201901c90506001818481612a5d57612a5d614f7f565b048201901c90506001818481612a7557612a75614f7f565b048201901c90506001818481612a8d57612a8d614f7f565b048201901c90506001818481612aa557612aa5614f7f565b048201901c9050610d3281828581612abf57612abf614f7f565b046139dd565b5f5b81831015612b1c575f612ada84846139eb565b5f8781526020902090915065ffffffffffff86169082015465ffffffffffff161115612b0857809250612b16565b612b13816001614e4e565b93505b50612ac7565b509392505050565b80515f906034811015612b3b576001915050610bce565b8281017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffec01517fffffffffffffffffffffffff000000000000000000000000000000000000000081167f2370726f706f7365723d3078000000000000000000000000000000000000000014612bb557600192505050610bce565b5f80612bc2602885614dc6565b90505b83811015612c50575f80612c10888481518110612be457612be4614dd9565b01602001517fff0000000000000000000000000000000000000000000000000000000000000016613a05565b9150915081612c285760019650505050505050610bce565b8060ff166004856001600160a01b0316901b179350505080612c4990614e17565b9050612bc5565b50856001600160a01b0316816001600160a01b031614935050505092915050565b5f612c858686868680519060200120611260565b905084518651141580612c9a57508351865114155b80612ca457508551155b15612cf2578551845186516040517f447b05d000000000000000000000000000000000000000000000000000000000815260048101939093526024830191909152604482015260640161134c565b5f81815260046020526040902054600160a01b900465ffffffffffff1615612d545780612d1e82610d39565b6040517f31b75e4d00000000000000000000000000000000000000000000000000000000815261134c9291905f90600401614fb2565b5f612d66600d5465ffffffffffff1690565b612d6e6110b9565b65ffffffffffff16612d809190614e4e565b90505f612d9d600d5463ffffffff66010000000000009091041690565b5f84815260046020526040902080547fffffffffffffffffffffffff0000000000000000000000000000000000000000166001600160a01b038716178155909150612de783611c63565b815465ffffffffffff91909116600160a01b027fffffffffffff000000000000ffffffffffffffffffffffffffffffffffffffff909116178155612e2a82613aef565b815463ffffffff919091167a010000000000000000000000000000000000000000000000000000027fffff00000000ffffffffffffffffffffffffffffffffffffffffffffffffffff90911617815588517f7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e090859087908c908c9067ffffffffffffffff811115612ebd57612ebd6143a8565b604051908082528060200260200182016040528015612ef057816020015b6060815260200190600190039081612edb5790505b508c89612efd8a82614e4e565b8e604051612f1399989796959493929190614ff3565b60405180910390a150505095945050505050565b606060ff8314612f4157612f3a83613b1f565b9050610bce565b818054612f4d90614d11565b80601f0160208091040260200160405190810160405280929190818152602001828054612f7990614d11565b8015612fc45780601f10612f9b57610100808354040283529160200191612fc4565b820191905f5260205f20905b815481529060010190602001808311612fa757829003601f168201915b50505050509050610bce565b5f610e80612fe6856001600160a01b031661144a565b8484613b5c565b5f81815260046020526040812054613038907a010000000000000000000000000000000000000000000000000000810463ffffffff1690600160a01b900465ffffffffffff16614fd4565b65ffffffffffff1692915050565b5f8183116130545781610d32565b5090919050565b80511561306b5780518082602001fd5b6040517f1425ea4200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b82545f90819080156131db575f6130b987611d19600185614dc6565b60408051808201909152905465ffffffffffff80821680845266010000000000009092046001600160d01b031660208401529192509087161015613129576040517f2520601d00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b805165ffffffffffff808816911603613178578461314c88611d19600186614dc6565b80546001600160d01b039290921666010000000000000265ffffffffffff9092169190911790556131cb565b6040805180820190915265ffffffffffff80881682526001600160d01b0380881660208085019182528b54600181018d555f8d815291909120945191519092166601000000000000029216919091179101555b6020015192508391506124a29050565b50506040805180820190915265ffffffffffff80851682526001600160d01b0380851660208085019182528854600181018a555f8a81529182209551925190931666010000000000000291909316179201919091559050816124a2565b5f8060095f9054906101000a90046001600160a01b03166001600160a01b031663f27a0c926040518163ffffffff1660e01b8152600401602060405180830381865afa15801561328a573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906132ae9190614f51565b90505f3060601b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001684186009546040517fb1c5f4270000000000000000000000000000000000000000000000000000000081529192506001600160a01b03169063b1c5f4279061332b908a908a908a905f9088906004016150c9565b602060405180830381865afa158015613346573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061336a9190614f51565b5f898152600a60205260408082209290925560095491517f8f2a0bb00000000000000000000000000000000000000000000000000000000081526001600160a01b0390921691638f2a0bb0916133cd918b918b918b919088908a90600401615116565b5f604051808303815f87803b1580156133e4575f80fd5b505af11580156133f6573d5f803e3d5ffd5b50505050610ea982426134099190614e4e565b611c63565b6009546001600160a01b031663e38335e5348686865f3060601b7fffffffffffffffffffffffffffffffffffffffff0000000000000000000000001688186040518763ffffffff1660e01b815260040161346c9594939291906150c9565b5f604051808303818588803b158015613483575f80fd5b505af1158015613495573d5f803e3d5ffd5b5050505f9687525050600a602052505060408320929092555050565b5f818152600760205260408120600181015481546134e2916001600160801b0390811691600160801b90041661516d565b6001600160801b0316613513610aa1855f9081526004602052604090205465ffffffffffff600160a01b9091041690565b11159392505050565b5f8061352a86868686613c08565b5f818152600a6020526040902054909150801561167d576009546040517fc4d252f5000000000000000000000000000000000000000000000000000000008152600481018390526001600160a01b039091169063c4d252f5906024015f604051808303815f87803b15801561359d575f80fd5b505af11580156135af573d5f803e3d5ffd5b5050505f838152600a602052604081205550509050949350505050565b5f306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561362457507f000000000000000000000000000000000000000000000000000000000000000046145b1561364e57507f000000000000000000000000000000000000000000000000000000000000000090565b610aeb604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b5f805f835160410361372c576020840151604085015160608601515f1a61371e88828585613ceb565b955095509550505050613737565b505081515f91506002905b9250925092565b5f805f856001600160a01b0316858560405160240161375e92919061518d565b60408051601f198184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f1626ba7e00000000000000000000000000000000000000000000000000000000179052516137c191906151a5565b5f60405180830381855afa9150503d805f81146137f9576040519150601f19603f3d011682016040523d82523d5f602084013e6137fe565b606091505b509150915081801561381257506020815110155b8015610fee575080517f1626ba7e00000000000000000000000000000000000000000000000000000000906138509083016020908101908401614f51565b149695505050505050565b5f61386a866115c06001612566565b505f86815260046020526040812054613894908790600160a01b900465ffffffffffff1685612fd0565b90506138a38787878487613db3565b82515f036138f757856001600160a01b03167fb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4888784886040516138ea94939291906151c0565b60405180910390a2610fee565b856001600160a01b03167fe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb871288878488886040516139389594939291906151e7565b60405180910390a29695505050505050565b5f80608083901c1561395e57608092831c92015b604083901c1561397057604092831c92015b602083901c1561398257602092831c92015b601083901c1561399457601092831c92015b600883901c156139a657600892831c92015b600483901c156139b857600492831c92015b600283901c156139ca57600292831c92015b600183901c15610bce5760010192915050565b5f8183106130545781610d32565b5f6139f96002848418614f93565b610d3290848416614e4e565b5f8060f883901c602f81118015613a1f5750603a8160ff16105b15613a52576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffd09091019350915050565b8060ff166040108015613a68575060478160ff16105b15613a9b576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffc99091019350915050565b8060ff166060108015613ab1575060678160ff16105b15613ae4576001947fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffa99091019350915050565b505f93849350915050565b5f63ffffffff821115611c95576040516306dfcc6560e41b8152602060048201526024810183905260440161134c565b60605f613b2b83613e88565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f7f00000000000000000000000000000000000000000000000000000000000000006040517f3a46b1a80000000000000000000000000000000000000000000000000000000081526001600160a01b038681166004830152602482018690529190911690633a46b1a890604401602060405180830381865afa158015613be4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e809190614f51565b5f80613c1686868686611260565b9050613c6481613c266007612566565b613c306006612566565b613c3a6002612566565b6001613c47600782615220565b613c52906002615319565b613c5c9190614dc6565b181818612588565b505f818152600460205260409081902080547effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f0100000000000000000000000000000000000000000000000000000000000000179055517f789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c9061181c9083815260200190565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613d2457505f91506003905082613da9565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613d75573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b038116613da057505f925060019150829050613da9565b92505f91508190505b9450945094915050565b815f03613dec576040517fbfc325cd00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8581526008602090815260408083206001600160a01b03881684529091529020546001600160801b03168211613e4f576040517f6d7d29ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f613e5983613ec8565b905081515f03613e7457613e6f86868387613efb565b613e80565b613e80868683856140bd565b505050505050565b5f60ff8216601f811115610bce576040517fb3512b0c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f6001600160801b03821115611c95576040516306dfcc6560e41b8152608060048201526024810183905260440161134c565b5f8481526008602090815260408083206001600160a01b03871684529091529020546001600160801b031615613f5d576040517f6d7d29ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8481526008602090815260408083206001600160a01b0387168452909152902080546fffffffffffffffffffffffffffffffff19166001600160801b03841617905560ff8116613ff7575f8481526007602052604081208054849290613fce9084906001600160801b031661516d565b92506101000a8154816001600160801b0302191690836001600160801b031602179055506140b7565b5f1960ff821601614032575f8481526007602052604090208054839190601090613fce908490600160801b90046001600160801b031661516d565b7ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe60ff821601614085575f8481526007602052604081206001018054849290613fce9084906001600160801b031661516d565b6040517f06b337c200000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b50505050565b80516030146140f8576040517fa653862c00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b6020818101516040808401515f888152600885528281206001600160a01b038916825290945290832054608083811c946001600160801b03948516949390911c929091169081836141498688614e4e565b6141539190614e4e565b61415d9190614e4e565b9050866001600160801b03168111156141a2576040517f6d7d29ea00000000000000000000000000000000000000000000000000000000815260040160405180910390fd5b5f8981526008602090815260408083206001600160a01b038c168452825280832080546fffffffffffffffffffffffffffffffff19166001600160801b03868116919091179091558c8452600783529281902081516060808201845282548087168352600160801b90048616948201949094526001919091015490931683820152805191820190528151819061423990899061516d565b6001600160801b03168152602001868360200151614257919061516d565b6001600160801b03168152602001858360400151614275919061516d565b6001600160801b039081169091525f9b8c526007602090815260409c8d90208351918401518316600160801b0291831691909117815591909b01516001909101805491909b166fffffffffffffffffffffffffffffffff1990911617909955505050505050505050565b5f602082840312156142ef575f80fd5b81357fffffffff0000000000000000000000000000000000000000000000000000000081168114610d32575f80fd5b5f6020828403121561432e575f80fd5b5035919050565b5f5b8381101561434f578181015183820152602001614337565b50505f910152565b5f815180845261436e816020860160208601614335565b601f01601f19169290920160200192915050565b602081525f610d326020830184614357565b6001600160a01b0381168114610be5575f80fd5b634e487b7160e01b5f52604160045260245ffd5b604051601f8201601f1916810167ffffffffffffffff811182821017156143e5576143e56143a8565b604052919050565b5f67ffffffffffffffff821115614406576144066143a8565b50601f01601f191660200190565b5f614426614421846143ed565b6143bc565b9050828152838383011115614439575f80fd5b828260208301375f602084830101529392505050565b5f82601f83011261445e575f80fd5b610d3283833560208501614414565b5f805f8060808587031215614480575f80fd5b843561448b81614394565b9350602085013561449b81614394565b925060408501359150606085013567ffffffffffffffff8111156144bd575f80fd5b6144c98782880161444f565b91505092959194509250565b5f67ffffffffffffffff8211156144ee576144ee6143a8565b5060051b60200190565b5f82601f830112614507575f80fd5b81356020614517614421836144d5565b82815260059290921b84018101918181019086841115614535575f80fd5b8286015b8481101561455957803561454c81614394565b8352918301918301614539565b509695505050505050565b5f82601f830112614573575f80fd5b81356020614583614421836144d5565b82815260059290921b840181019181810190868411156145a1575f80fd5b8286015b8481101561455957803583529183019183016145a5565b5f82601f8301126145cb575f80fd5b813560206145db614421836144d5565b82815260059290921b840181019181810190868411156145f9575f80fd5b8286015b8481101561455957803567ffffffffffffffff81111561461c575f8081fd5b61462a8986838b010161444f565b8452509183019183016145fd565b5f805f806080858703121561464b575f80fd5b843567ffffffffffffffff80821115614662575f80fd5b61466e888389016144f8565b95506020870135915080821115614683575f80fd5b61468f88838901614564565b945060408701359150808211156146a4575f80fd5b506146b1878288016145bc565b949793965093946060013593505050565b65ffffffffffff81168114610be5575f80fd5b5f602082840312156146e5575f80fd5b8135610d32816146c2565b5f8060408385031215614701575f80fd5b82359150602083013561471381614394565b809150509250929050565b634e487b7160e01b5f52602160045260245ffd5b6008811061474e57634e487b7160e01b5f52602160045260245ffd5b9052565b60208101610bce8284614732565b803560ff81168114610e5c575f80fd5b5f8060408385031215614781575f80fd5b8235915061479160208401614760565b90509250929050565b5f8083601f8401126147aa575f80fd5b50813567ffffffffffffffff8111156147c1575f80fd5b6020830191508360208285010111156147d8575f80fd5b9250929050565b5f805f805f805f60c0888a0312156147f5575f80fd5b8735965061480560208901614760565b9550604088013561481581614394565b9450606088013567ffffffffffffffff80821115614831575f80fd5b61483d8b838c0161479a565b909650945060808a0135915080821115614855575f80fd5b6148618b838c0161444f565b935060a08a0135915080821115614876575f80fd5b506148838a828b0161444f565b91505092959891949750929550565b5f805f805f608086880312156148a6575f80fd5b853594506148b660208701614760565b9350604086013567ffffffffffffffff808211156148d2575f80fd5b6148de89838a0161479a565b909550935060608801359150808211156148f6575f80fd5b506149038882890161444f565b9150509295509295909350565b5f805f8060608587031215614923575f80fd5b8435935061493360208601614760565b9250604085013567ffffffffffffffff81111561494e575f80fd5b61495a8782880161479a565b95989497509550505050565b5f805f8060808587031215614979575f80fd5b843567ffffffffffffffff80821115614990575f80fd5b61499c888389016144f8565b955060208701359150808211156149b1575f80fd5b6149bd88838901614564565b945060408701359150808211156149d2575f80fd5b6149de888389016145bc565b935060608701359150808211156149f3575f80fd5b508501601f81018713614a04575f80fd5b6144c987823560208401614414565b5f60208284031215614a23575f80fd5b8135610d3281614394565b5f8151808452602080850194508084015f5b83811015614a5c57815187529582019590820190600101614a40565b509495945050505050565b7fff000000000000000000000000000000000000000000000000000000000000008816815260e060208201525f614aa160e0830189614357565b8281036040840152614ab38189614357565b90508660608401526001600160a01b03861660808401528460a084015282810360c0840152614ae28185614a2e565b9a9950505050505050505050565b5f805f8060808587031215614b03575f80fd5b84359350614b1360208601614760565b92506040850135614b2381614394565b9150606085013567ffffffffffffffff8111156144bd575f80fd5b5f805f60608486031215614b50575f80fd5b8335614b5b81614394565b925060208401359150604084013567ffffffffffffffff811115614b7d575f80fd5b614b898682870161444f565b9150509250925092565b5f805f805f60a08688031215614ba7575f80fd5b8535614bb281614394565b94506020860135614bc281614394565b9350604086013567ffffffffffffffff80821115614bde575f80fd5b614bea89838a01614564565b94506060880135915080821115614bff575f80fd5b614c0b89838a01614564565b935060808801359150808211156148f6575f80fd5b5f805f8060608587031215614c33575f80fd5b8435614c3e81614394565b935060208501359250604085013567ffffffffffffffff81111561494e575f80fd5b5f60208284031215614c70575f80fd5b813563ffffffff81168114610d32575f80fd5b5f8060408385031215614c94575f80fd5b8235614c9f81614394565b946020939093013593505050565b5f805f805f60a08688031215614cc1575f80fd5b8535614ccc81614394565b94506020860135614cdc81614394565b93506040860135925060608601359150608086013567ffffffffffffffff811115614d05575f80fd5b6149038882890161444f565b600181811c90821680614d2557607f821691505b602082108103614d4357634e487b7160e01b5f52602260045260245ffd5b50919050565b5f60208284031215614d59575f80fd5b815167ffffffffffffffff811115614d6f575f80fd5b8201601f81018413614d7f575f80fd5b8051614d8d614421826143ed565b818152856020838501011115614da1575f80fd5b610ca7826020830160208601614335565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610bce57610bce614db2565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215614dfd575f80fd5b8151610d32816146c2565b818382375f9101908152919050565b5f5f198203614e2857614e28614db2565b5060010190565b5f60208284031215614e3f575f80fd5b81518015158114610d32575f80fd5b80820180821115610bce57610bce614db2565b65ffffffffffff82811682821603908082111561159b5761159b614db2565b5f8151808452602080850194508084015f5b83811015614a5c5781516001600160a01b031687529582019590820190600101614e92565b5f815180845260208085019450848260051b86018286015f5b85811015614efa578383038952614ee8838351614357565b98850198925090840190600101614ed0565b5090979650505050505050565b608081525f614f196080830187614e80565b8281036020840152614f2b8187614a2e565b90508281036040840152614f3f8186614eb7565b91505082606083015295945050505050565b5f60208284031215614f61575f80fd5b5051919050565b8082028115828204841417610bce57610bce614db2565b634e487b7160e01b5f52601260045260245ffd5b5f82614fad57634e487b7160e01b5f52601260045260245ffd5b500490565b83815260608101614fc66020830185614732565b826040830152949350505050565b65ffffffffffff81811683821601908082111561159b5761159b614db2565b5f6101208b835260206001600160a01b038c168185015281604085015261501c8285018c614e80565b91508382036060850152615030828b614a2e565b915083820360808501528189518084528284019150828160051b850101838c015f5b8381101561508057601f1987840301855261506e838351614357565b94860194925090850190600101615052565b505086810360a0880152615094818c614eb7565b9450505050508560c08401528460e08401528281036101008401526150b98185614357565b9c9b505050505050505050505050565b60a081525f6150db60a0830188614e80565b82810360208401526150ed8188614a2e565b905082810360408401526151018187614eb7565b60608401959095525050608001529392505050565b60c081525f61512860c0830189614e80565b828103602084015261513a8189614a2e565b9050828103604084015261514e8188614eb7565b60608401969096525050608081019290925260a0909101529392505050565b6001600160801b0381811683821601908082111561159b5761159b614db2565b828152604060208201525f610e806040830184614357565b5f82516151b6818460208701614335565b9190910192915050565b84815260ff84166020820152826040820152608060608201525f610fee6080830184614357565b85815260ff8516602082015283604082015260a060608201525f61520e60a0830185614357565b8281036080840152610ea98185614357565b60ff8181168382160190811115610bce57610bce614db2565b600181815b8085111561527357815f190482111561525957615259614db2565b8085161561526657918102915b93841c939080029061523e565b509250929050565b5f8261528957506001610bce565b8161529557505f610bce565b81600181146152ab57600281146152b5576152d1565b6001915050610bce565b60ff8411156152c6576152c6614db2565b50506001821b610bce565b5060208310610133831016604e8410600b84101617156152f4575081810a610bce565b6152fe8383615239565b805f190482111561531157615311614db2565b029392505050565b5f610d3260ff84168361527b56fe737570706f72743d627261766f2671756f72756d3d666f722c6162737461696e26706172616d733d6672616374696f6e616ca26469706673582212200fcac9b29aae8a870610544b669cd6e1c25a5e2694ebfa90dde6f66b19f5ccbf64736f6c63430008150033",
}

// GlifGovernorABI is the input ABI used to generate the binding from.
// Deprecated: Use GlifGovernorMetaData.ABI instead.
var GlifGovernorABI = GlifGovernorMetaData.ABI

// GlifGovernorBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use GlifGovernorMetaData.Bin instead.
var GlifGovernorBin = GlifGovernorMetaData.Bin

// DeployGlifGovernor deploys a new Ethereum contract, binding an instance of GlifGovernor to it.
func DeployGlifGovernor(auth *bind.TransactOpts, backend bind.ContractBackend, _token common.Address, _initialQuorumNumerator *big.Int, _initialVotingDelay *big.Int, _initialVotingPeriod uint32, _initialProposalThreshold *big.Int, _initialVoteExtension *big.Int, _timelock common.Address) (common.Address, *types.Transaction, *GlifGovernor, error) {
	parsed, err := GlifGovernorMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(GlifGovernorBin), backend, _token, _initialQuorumNumerator, _initialVotingDelay, _initialVotingPeriod, _initialProposalThreshold, _initialVoteExtension, _timelock)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &GlifGovernor{GlifGovernorCaller: GlifGovernorCaller{contract: contract}, GlifGovernorTransactor: GlifGovernorTransactor{contract: contract}, GlifGovernorFilterer: GlifGovernorFilterer{contract: contract}}, nil
}

// GlifGovernor is an auto generated Go binding around an Ethereum contract.
type GlifGovernor struct {
	GlifGovernorCaller     // Read-only binding to the contract
	GlifGovernorTransactor // Write-only binding to the contract
	GlifGovernorFilterer   // Log filterer for contract events
}

// GlifGovernorCaller is an auto generated read-only Go binding around an Ethereum contract.
type GlifGovernorCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlifGovernorTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GlifGovernorTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlifGovernorFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GlifGovernorFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GlifGovernorSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GlifGovernorSession struct {
	Contract     *GlifGovernor     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GlifGovernorCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GlifGovernorCallerSession struct {
	Contract *GlifGovernorCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// GlifGovernorTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GlifGovernorTransactorSession struct {
	Contract     *GlifGovernorTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// GlifGovernorRaw is an auto generated low-level Go binding around an Ethereum contract.
type GlifGovernorRaw struct {
	Contract *GlifGovernor // Generic contract binding to access the raw methods on
}

// GlifGovernorCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GlifGovernorCallerRaw struct {
	Contract *GlifGovernorCaller // Generic read-only contract binding to access the raw methods on
}

// GlifGovernorTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GlifGovernorTransactorRaw struct {
	Contract *GlifGovernorTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGlifGovernor creates a new instance of GlifGovernor, bound to a specific deployed contract.
func NewGlifGovernor(address common.Address, backend bind.ContractBackend) (*GlifGovernor, error) {
	contract, err := bindGlifGovernor(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GlifGovernor{GlifGovernorCaller: GlifGovernorCaller{contract: contract}, GlifGovernorTransactor: GlifGovernorTransactor{contract: contract}, GlifGovernorFilterer: GlifGovernorFilterer{contract: contract}}, nil
}

// NewGlifGovernorCaller creates a new read-only instance of GlifGovernor, bound to a specific deployed contract.
func NewGlifGovernorCaller(address common.Address, caller bind.ContractCaller) (*GlifGovernorCaller, error) {
	contract, err := bindGlifGovernor(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GlifGovernorCaller{contract: contract}, nil
}

// NewGlifGovernorTransactor creates a new write-only instance of GlifGovernor, bound to a specific deployed contract.
func NewGlifGovernorTransactor(address common.Address, transactor bind.ContractTransactor) (*GlifGovernorTransactor, error) {
	contract, err := bindGlifGovernor(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GlifGovernorTransactor{contract: contract}, nil
}

// NewGlifGovernorFilterer creates a new log filterer instance of GlifGovernor, bound to a specific deployed contract.
func NewGlifGovernorFilterer(address common.Address, filterer bind.ContractFilterer) (*GlifGovernorFilterer, error) {
	contract, err := bindGlifGovernor(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GlifGovernorFilterer{contract: contract}, nil
}

// bindGlifGovernor binds a generic wrapper to an already deployed contract.
func bindGlifGovernor(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := GlifGovernorMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlifGovernor *GlifGovernorRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlifGovernor.Contract.GlifGovernorCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlifGovernor *GlifGovernorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlifGovernor.Contract.GlifGovernorTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlifGovernor *GlifGovernorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlifGovernor.Contract.GlifGovernorTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GlifGovernor *GlifGovernorCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GlifGovernor.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GlifGovernor *GlifGovernorTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlifGovernor.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GlifGovernor *GlifGovernorTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GlifGovernor.Contract.contract.Transact(opts, method, params...)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_GlifGovernor *GlifGovernorCaller) BALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_GlifGovernor *GlifGovernorSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _GlifGovernor.Contract.BALLOTTYPEHASH(&_GlifGovernor.CallOpts)
}

// BALLOTTYPEHASH is a free data retrieval call binding the contract method 0xdeaaa7cc.
//
// Solidity: function BALLOT_TYPEHASH() view returns(bytes32)
func (_GlifGovernor *GlifGovernorCallerSession) BALLOTTYPEHASH() ([32]byte, error) {
	return _GlifGovernor.Contract.BALLOTTYPEHASH(&_GlifGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_GlifGovernor *GlifGovernorCaller) CLOCKMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "CLOCK_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_GlifGovernor *GlifGovernorSession) CLOCKMODE() (string, error) {
	return _GlifGovernor.Contract.CLOCKMODE(&_GlifGovernor.CallOpts)
}

// CLOCKMODE is a free data retrieval call binding the contract method 0x4bf5d7e9.
//
// Solidity: function CLOCK_MODE() view returns(string)
func (_GlifGovernor *GlifGovernorCallerSession) CLOCKMODE() (string, error) {
	return _GlifGovernor.Contract.CLOCKMODE(&_GlifGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_GlifGovernor *GlifGovernorCaller) COUNTINGMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "COUNTING_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_GlifGovernor *GlifGovernorSession) COUNTINGMODE() (string, error) {
	return _GlifGovernor.Contract.COUNTINGMODE(&_GlifGovernor.CallOpts)
}

// COUNTINGMODE is a free data retrieval call binding the contract method 0xdd4e2ba5.
//
// Solidity: function COUNTING_MODE() pure returns(string)
func (_GlifGovernor *GlifGovernorCallerSession) COUNTINGMODE() (string, error) {
	return _GlifGovernor.Contract.COUNTINGMODE(&_GlifGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_GlifGovernor *GlifGovernorCaller) EXTENDEDBALLOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "EXTENDED_BALLOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_GlifGovernor *GlifGovernorSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _GlifGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_GlifGovernor.CallOpts)
}

// EXTENDEDBALLOTTYPEHASH is a free data retrieval call binding the contract method 0x2fe3e261.
//
// Solidity: function EXTENDED_BALLOT_TYPEHASH() view returns(bytes32)
func (_GlifGovernor *GlifGovernorCallerSession) EXTENDEDBALLOTTYPEHASH() ([32]byte, error) {
	return _GlifGovernor.Contract.EXTENDEDBALLOTTYPEHASH(&_GlifGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_GlifGovernor *GlifGovernorCaller) Clock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "clock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_GlifGovernor *GlifGovernorSession) Clock() (*big.Int, error) {
	return _GlifGovernor.Contract.Clock(&_GlifGovernor.CallOpts)
}

// Clock is a free data retrieval call binding the contract method 0x91ddadf4.
//
// Solidity: function clock() view returns(uint48)
func (_GlifGovernor *GlifGovernorCallerSession) Clock() (*big.Int, error) {
	return _GlifGovernor.Contract.Clock(&_GlifGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GlifGovernor *GlifGovernorCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GlifGovernor *GlifGovernorSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GlifGovernor.Contract.Eip712Domain(&_GlifGovernor.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_GlifGovernor *GlifGovernorCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _GlifGovernor.Contract.Eip712Domain(&_GlifGovernor.CallOpts)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) GetVotes(opts *bind.CallOpts, account common.Address, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "getVotes", account, timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.GetVotes(&_GlifGovernor.CallOpts, account, timepoint)
}

// GetVotes is a free data retrieval call binding the contract method 0xeb9019d4.
//
// Solidity: function getVotes(address account, uint256 timepoint) view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) GetVotes(account common.Address, timepoint *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.GetVotes(&_GlifGovernor.CallOpts, account, timepoint)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) GetVotesWithParams(opts *bind.CallOpts, account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "getVotesWithParams", account, timepoint, params)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _GlifGovernor.Contract.GetVotesWithParams(&_GlifGovernor.CallOpts, account, timepoint, params)
}

// GetVotesWithParams is a free data retrieval call binding the contract method 0x9a802a6d.
//
// Solidity: function getVotesWithParams(address account, uint256 timepoint, bytes params) view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) GetVotesWithParams(account common.Address, timepoint *big.Int, params []byte) (*big.Int, error) {
	return _GlifGovernor.Contract.GetVotesWithParams(&_GlifGovernor.CallOpts, account, timepoint, params)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_GlifGovernor *GlifGovernorCaller) HasVoted(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (bool, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "hasVoted", proposalId, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_GlifGovernor *GlifGovernorSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _GlifGovernor.Contract.HasVoted(&_GlifGovernor.CallOpts, proposalId, account)
}

// HasVoted is a free data retrieval call binding the contract method 0x43859632.
//
// Solidity: function hasVoted(uint256 proposalId, address account) view returns(bool)
func (_GlifGovernor *GlifGovernorCallerSession) HasVoted(proposalId *big.Int, account common.Address) (bool, error) {
	return _GlifGovernor.Contract.HasVoted(&_GlifGovernor.CallOpts, proposalId, account)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) HashProposal(opts *bind.CallOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "hashProposal", targets, values, calldatas, descriptionHash)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_GlifGovernor *GlifGovernorSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _GlifGovernor.Contract.HashProposal(&_GlifGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// HashProposal is a free data retrieval call binding the contract method 0xc59057e4.
//
// Solidity: function hashProposal(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) pure returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) HashProposal(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*big.Int, error) {
	return _GlifGovernor.Contract.HashProposal(&_GlifGovernor.CallOpts, targets, values, calldatas, descriptionHash)
}

// LateQuorumVoteExtension is a free data retrieval call binding the contract method 0x32b8113e.
//
// Solidity: function lateQuorumVoteExtension() view returns(uint48)
func (_GlifGovernor *GlifGovernorCaller) LateQuorumVoteExtension(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "lateQuorumVoteExtension")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LateQuorumVoteExtension is a free data retrieval call binding the contract method 0x32b8113e.
//
// Solidity: function lateQuorumVoteExtension() view returns(uint48)
func (_GlifGovernor *GlifGovernorSession) LateQuorumVoteExtension() (*big.Int, error) {
	return _GlifGovernor.Contract.LateQuorumVoteExtension(&_GlifGovernor.CallOpts)
}

// LateQuorumVoteExtension is a free data retrieval call binding the contract method 0x32b8113e.
//
// Solidity: function lateQuorumVoteExtension() view returns(uint48)
func (_GlifGovernor *GlifGovernorCallerSession) LateQuorumVoteExtension() (*big.Int, error) {
	return _GlifGovernor.Contract.LateQuorumVoteExtension(&_GlifGovernor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GlifGovernor *GlifGovernorCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GlifGovernor *GlifGovernorSession) Name() (string, error) {
	return _GlifGovernor.Contract.Name(&_GlifGovernor.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_GlifGovernor *GlifGovernorCallerSession) Name() (string, error) {
	return _GlifGovernor.Contract.Name(&_GlifGovernor.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) Nonces(owner common.Address) (*big.Int, error) {
	return _GlifGovernor.Contract.Nonces(&_GlifGovernor.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _GlifGovernor.Contract.Nonces(&_GlifGovernor.CallOpts, owner)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) ProposalDeadline(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "proposalDeadline", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.ProposalDeadline(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalDeadline is a free data retrieval call binding the contract method 0xc01f9e37.
//
// Solidity: function proposalDeadline(uint256 proposalId) view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) ProposalDeadline(proposalId *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.ProposalDeadline(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) ProposalEta(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "proposalEta", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.ProposalEta(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalEta is a free data retrieval call binding the contract method 0xab58fb8e.
//
// Solidity: function proposalEta(uint256 proposalId) view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) ProposalEta(proposalId *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.ProposalEta(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_GlifGovernor *GlifGovernorCaller) ProposalNeedsQueuing(opts *bind.CallOpts, proposalId *big.Int) (bool, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "proposalNeedsQueuing", proposalId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_GlifGovernor *GlifGovernorSession) ProposalNeedsQueuing(proposalId *big.Int) (bool, error) {
	return _GlifGovernor.Contract.ProposalNeedsQueuing(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalNeedsQueuing is a free data retrieval call binding the contract method 0xa9a95294.
//
// Solidity: function proposalNeedsQueuing(uint256 proposalId) view returns(bool)
func (_GlifGovernor *GlifGovernorCallerSession) ProposalNeedsQueuing(proposalId *big.Int) (bool, error) {
	return _GlifGovernor.Contract.ProposalNeedsQueuing(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_GlifGovernor *GlifGovernorCaller) ProposalProposer(opts *bind.CallOpts, proposalId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "proposalProposer", proposalId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_GlifGovernor *GlifGovernorSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _GlifGovernor.Contract.ProposalProposer(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalProposer is a free data retrieval call binding the contract method 0x143489d0.
//
// Solidity: function proposalProposer(uint256 proposalId) view returns(address)
func (_GlifGovernor *GlifGovernorCallerSession) ProposalProposer(proposalId *big.Int) (common.Address, error) {
	return _GlifGovernor.Contract.ProposalProposer(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) ProposalSnapshot(opts *bind.CallOpts, proposalId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "proposalSnapshot", proposalId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.ProposalSnapshot(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalSnapshot is a free data retrieval call binding the contract method 0x2d63f693.
//
// Solidity: function proposalSnapshot(uint256 proposalId) view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) ProposalSnapshot(proposalId *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.ProposalSnapshot(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) ProposalThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "proposalThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) ProposalThreshold() (*big.Int, error) {
	return _GlifGovernor.Contract.ProposalThreshold(&_GlifGovernor.CallOpts)
}

// ProposalThreshold is a free data retrieval call binding the contract method 0xb58131b0.
//
// Solidity: function proposalThreshold() view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) ProposalThreshold() (*big.Int, error) {
	return _GlifGovernor.Contract.ProposalThreshold(&_GlifGovernor.CallOpts)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_GlifGovernor *GlifGovernorCaller) ProposalVotes(opts *bind.CallOpts, proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "proposalVotes", proposalId)

	outstruct := new(struct {
		AgainstVotes *big.Int
		ForVotes     *big.Int
		AbstainVotes *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AgainstVotes = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ForVotes = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AbstainVotes = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_GlifGovernor *GlifGovernorSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _GlifGovernor.Contract.ProposalVotes(&_GlifGovernor.CallOpts, proposalId)
}

// ProposalVotes is a free data retrieval call binding the contract method 0x544ffc9c.
//
// Solidity: function proposalVotes(uint256 proposalId) view returns(uint256 againstVotes, uint256 forVotes, uint256 abstainVotes)
func (_GlifGovernor *GlifGovernorCallerSession) ProposalVotes(proposalId *big.Int) (struct {
	AgainstVotes *big.Int
	ForVotes     *big.Int
	AbstainVotes *big.Int
}, error) {
	return _GlifGovernor.Contract.ProposalVotes(&_GlifGovernor.CallOpts, proposalId)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) Quorum(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "quorum", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) Quorum(timepoint *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.Quorum(&_GlifGovernor.CallOpts, timepoint)
}

// Quorum is a free data retrieval call binding the contract method 0xf8ce560a.
//
// Solidity: function quorum(uint256 timepoint) view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) Quorum(timepoint *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.Quorum(&_GlifGovernor.CallOpts, timepoint)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) QuorumDenominator(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "quorumDenominator")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) QuorumDenominator() (*big.Int, error) {
	return _GlifGovernor.Contract.QuorumDenominator(&_GlifGovernor.CallOpts)
}

// QuorumDenominator is a free data retrieval call binding the contract method 0x97c3d334.
//
// Solidity: function quorumDenominator() view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) QuorumDenominator() (*big.Int, error) {
	return _GlifGovernor.Contract.QuorumDenominator(&_GlifGovernor.CallOpts)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) QuorumNumerator(opts *bind.CallOpts, timepoint *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "quorumNumerator", timepoint)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.QuorumNumerator(&_GlifGovernor.CallOpts, timepoint)
}

// QuorumNumerator is a free data retrieval call binding the contract method 0x60c4247f.
//
// Solidity: function quorumNumerator(uint256 timepoint) view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) QuorumNumerator(timepoint *big.Int) (*big.Int, error) {
	return _GlifGovernor.Contract.QuorumNumerator(&_GlifGovernor.CallOpts, timepoint)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) QuorumNumerator0(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "quorumNumerator0")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) QuorumNumerator0() (*big.Int, error) {
	return _GlifGovernor.Contract.QuorumNumerator0(&_GlifGovernor.CallOpts)
}

// QuorumNumerator0 is a free data retrieval call binding the contract method 0xa7713a70.
//
// Solidity: function quorumNumerator() view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) QuorumNumerator0() (*big.Int, error) {
	return _GlifGovernor.Contract.QuorumNumerator0(&_GlifGovernor.CallOpts)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_GlifGovernor *GlifGovernorCaller) State(opts *bind.CallOpts, proposalId *big.Int) (uint8, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "state", proposalId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_GlifGovernor *GlifGovernorSession) State(proposalId *big.Int) (uint8, error) {
	return _GlifGovernor.Contract.State(&_GlifGovernor.CallOpts, proposalId)
}

// State is a free data retrieval call binding the contract method 0x3e4f49e6.
//
// Solidity: function state(uint256 proposalId) view returns(uint8)
func (_GlifGovernor *GlifGovernorCallerSession) State(proposalId *big.Int) (uint8, error) {
	return _GlifGovernor.Contract.State(&_GlifGovernor.CallOpts, proposalId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GlifGovernor *GlifGovernorCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GlifGovernor *GlifGovernorSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GlifGovernor.Contract.SupportsInterface(&_GlifGovernor.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_GlifGovernor *GlifGovernorCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _GlifGovernor.Contract.SupportsInterface(&_GlifGovernor.CallOpts, interfaceId)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_GlifGovernor *GlifGovernorCaller) Timelock(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "timelock")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_GlifGovernor *GlifGovernorSession) Timelock() (common.Address, error) {
	return _GlifGovernor.Contract.Timelock(&_GlifGovernor.CallOpts)
}

// Timelock is a free data retrieval call binding the contract method 0xd33219b4.
//
// Solidity: function timelock() view returns(address)
func (_GlifGovernor *GlifGovernorCallerSession) Timelock() (common.Address, error) {
	return _GlifGovernor.Contract.Timelock(&_GlifGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GlifGovernor *GlifGovernorCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GlifGovernor *GlifGovernorSession) Token() (common.Address, error) {
	return _GlifGovernor.Contract.Token(&_GlifGovernor.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_GlifGovernor *GlifGovernorCallerSession) Token() (common.Address, error) {
	return _GlifGovernor.Contract.Token(&_GlifGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_GlifGovernor *GlifGovernorCaller) Version(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_GlifGovernor *GlifGovernorSession) Version() (string, error) {
	return _GlifGovernor.Contract.Version(&_GlifGovernor.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(string)
func (_GlifGovernor *GlifGovernorCallerSession) Version() (string, error) {
	return _GlifGovernor.Contract.Version(&_GlifGovernor.CallOpts)
}

// VoteWeightCast is a free data retrieval call binding the contract method 0x342cfab6.
//
// Solidity: function voteWeightCast(uint256 proposalId, address account) view returns(uint128)
func (_GlifGovernor *GlifGovernorCaller) VoteWeightCast(opts *bind.CallOpts, proposalId *big.Int, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "voteWeightCast", proposalId, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VoteWeightCast is a free data retrieval call binding the contract method 0x342cfab6.
//
// Solidity: function voteWeightCast(uint256 proposalId, address account) view returns(uint128)
func (_GlifGovernor *GlifGovernorSession) VoteWeightCast(proposalId *big.Int, account common.Address) (*big.Int, error) {
	return _GlifGovernor.Contract.VoteWeightCast(&_GlifGovernor.CallOpts, proposalId, account)
}

// VoteWeightCast is a free data retrieval call binding the contract method 0x342cfab6.
//
// Solidity: function voteWeightCast(uint256 proposalId, address account) view returns(uint128)
func (_GlifGovernor *GlifGovernorCallerSession) VoteWeightCast(proposalId *big.Int, account common.Address) (*big.Int, error) {
	return _GlifGovernor.Contract.VoteWeightCast(&_GlifGovernor.CallOpts, proposalId, account)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) VotingDelay(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "votingDelay")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) VotingDelay() (*big.Int, error) {
	return _GlifGovernor.Contract.VotingDelay(&_GlifGovernor.CallOpts)
}

// VotingDelay is a free data retrieval call binding the contract method 0x3932abb1.
//
// Solidity: function votingDelay() view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) VotingDelay() (*big.Int, error) {
	return _GlifGovernor.Contract.VotingDelay(&_GlifGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_GlifGovernor *GlifGovernorCaller) VotingPeriod(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _GlifGovernor.contract.Call(opts, &out, "votingPeriod")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_GlifGovernor *GlifGovernorSession) VotingPeriod() (*big.Int, error) {
	return _GlifGovernor.Contract.VotingPeriod(&_GlifGovernor.CallOpts)
}

// VotingPeriod is a free data retrieval call binding the contract method 0x02a251a3.
//
// Solidity: function votingPeriod() view returns(uint256)
func (_GlifGovernor *GlifGovernorCallerSession) VotingPeriod() (*big.Int, error) {
	return _GlifGovernor.Contract.VotingPeriod(&_GlifGovernor.CallOpts)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactor) Cancel(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "cancel", targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GlifGovernor *GlifGovernorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.Cancel(&_GlifGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Cancel is a paid mutator transaction binding the contract method 0x452115d6.
//
// Solidity: function cancel(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactorSession) Cancel(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.Cancel(&_GlifGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactor) CastVote(opts *bind.TransactOpts, proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "castVote", proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_GlifGovernor *GlifGovernorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _GlifGovernor.Contract.CastVote(&_GlifGovernor.TransactOpts, proposalId, support)
}

// CastVote is a paid mutator transaction binding the contract method 0x56781388.
//
// Solidity: function castVote(uint256 proposalId, uint8 support) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactorSession) CastVote(proposalId *big.Int, support uint8) (*types.Transaction, error) {
	return _GlifGovernor.Contract.CastVote(&_GlifGovernor.TransactOpts, proposalId, support)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactor) CastVoteBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "castVoteBySig", proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_GlifGovernor *GlifGovernorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.CastVoteBySig(&_GlifGovernor.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteBySig is a paid mutator transaction binding the contract method 0x8ff262e3.
//
// Solidity: function castVoteBySig(uint256 proposalId, uint8 support, address voter, bytes signature) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactorSession) CastVoteBySig(proposalId *big.Int, support uint8, voter common.Address, signature []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.CastVoteBySig(&_GlifGovernor.TransactOpts, proposalId, support, voter, signature)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactor) CastVoteWithReason(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "castVoteWithReason", proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_GlifGovernor *GlifGovernorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _GlifGovernor.Contract.CastVoteWithReason(&_GlifGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReason is a paid mutator transaction binding the contract method 0x7b3c71d3.
//
// Solidity: function castVoteWithReason(uint256 proposalId, uint8 support, string reason) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactorSession) CastVoteWithReason(proposalId *big.Int, support uint8, reason string) (*types.Transaction, error) {
	return _GlifGovernor.Contract.CastVoteWithReason(&_GlifGovernor.TransactOpts, proposalId, support, reason)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactor) CastVoteWithReasonAndParams(opts *bind.TransactOpts, proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "castVoteWithReasonAndParams", proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_GlifGovernor *GlifGovernorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.CastVoteWithReasonAndParams(&_GlifGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParams is a paid mutator transaction binding the contract method 0x5f398a14.
//
// Solidity: function castVoteWithReasonAndParams(uint256 proposalId, uint8 support, string reason, bytes params) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactorSession) CastVoteWithReasonAndParams(proposalId *big.Int, support uint8, reason string, params []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.CastVoteWithReasonAndParams(&_GlifGovernor.TransactOpts, proposalId, support, reason, params)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactor) CastVoteWithReasonAndParamsBySig(opts *bind.TransactOpts, proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "castVoteWithReasonAndParamsBySig", proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_GlifGovernor *GlifGovernorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_GlifGovernor.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// CastVoteWithReasonAndParamsBySig is a paid mutator transaction binding the contract method 0x5b8d0e0d.
//
// Solidity: function castVoteWithReasonAndParamsBySig(uint256 proposalId, uint8 support, address voter, string reason, bytes params, bytes signature) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactorSession) CastVoteWithReasonAndParamsBySig(proposalId *big.Int, support uint8, voter common.Address, reason string, params []byte, signature []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.CastVoteWithReasonAndParamsBySig(&_GlifGovernor.TransactOpts, proposalId, support, voter, reason, params, signature)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_GlifGovernor *GlifGovernorTransactor) Execute(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "execute", targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_GlifGovernor *GlifGovernorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.Execute(&_GlifGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Execute is a paid mutator transaction binding the contract method 0x2656227d.
//
// Solidity: function execute(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) payable returns(uint256)
func (_GlifGovernor *GlifGovernorTransactorSession) Execute(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.Execute(&_GlifGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_GlifGovernor *GlifGovernorTransactor) OnERC1155BatchReceived(opts *bind.TransactOpts, operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "onERC1155BatchReceived", operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_GlifGovernor *GlifGovernorSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.OnERC1155BatchReceived(&_GlifGovernor.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a paid mutator transaction binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) returns(bytes4)
func (_GlifGovernor *GlifGovernorTransactorSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.OnERC1155BatchReceived(&_GlifGovernor.TransactOpts, operator, from, ids, values, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_GlifGovernor *GlifGovernorTransactor) OnERC1155Received(opts *bind.TransactOpts, operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "onERC1155Received", operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_GlifGovernor *GlifGovernorSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.OnERC1155Received(&_GlifGovernor.TransactOpts, operator, from, id, value, data)
}

// OnERC1155Received is a paid mutator transaction binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) returns(bytes4)
func (_GlifGovernor *GlifGovernorTransactorSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.OnERC1155Received(&_GlifGovernor.TransactOpts, operator, from, id, value, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_GlifGovernor *GlifGovernorTransactor) OnERC721Received(opts *bind.TransactOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "onERC721Received", operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_GlifGovernor *GlifGovernorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.OnERC721Received(&_GlifGovernor.TransactOpts, operator, from, tokenId, data)
}

// OnERC721Received is a paid mutator transaction binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) returns(bytes4)
func (_GlifGovernor *GlifGovernorTransactorSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.OnERC721Received(&_GlifGovernor.TransactOpts, operator, from, tokenId, data)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactor) Propose(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "propose", targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_GlifGovernor *GlifGovernorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _GlifGovernor.Contract.Propose(&_GlifGovernor.TransactOpts, targets, values, calldatas, description)
}

// Propose is a paid mutator transaction binding the contract method 0x7d5e81e2.
//
// Solidity: function propose(address[] targets, uint256[] values, bytes[] calldatas, string description) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactorSession) Propose(targets []common.Address, values []*big.Int, calldatas [][]byte, description string) (*types.Transaction, error) {
	return _GlifGovernor.Contract.Propose(&_GlifGovernor.TransactOpts, targets, values, calldatas, description)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactor) Queue(opts *bind.TransactOpts, targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "queue", targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GlifGovernor *GlifGovernorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.Queue(&_GlifGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Queue is a paid mutator transaction binding the contract method 0x160cbed7.
//
// Solidity: function queue(address[] targets, uint256[] values, bytes[] calldatas, bytes32 descriptionHash) returns(uint256)
func (_GlifGovernor *GlifGovernorTransactorSession) Queue(targets []common.Address, values []*big.Int, calldatas [][]byte, descriptionHash [32]byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.Queue(&_GlifGovernor.TransactOpts, targets, values, calldatas, descriptionHash)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_GlifGovernor *GlifGovernorTransactor) Relay(opts *bind.TransactOpts, target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "relay", target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_GlifGovernor *GlifGovernorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.Relay(&_GlifGovernor.TransactOpts, target, value, data)
}

// Relay is a paid mutator transaction binding the contract method 0xc28bc2fa.
//
// Solidity: function relay(address target, uint256 value, bytes data) payable returns()
func (_GlifGovernor *GlifGovernorTransactorSession) Relay(target common.Address, value *big.Int, data []byte) (*types.Transaction, error) {
	return _GlifGovernor.Contract.Relay(&_GlifGovernor.TransactOpts, target, value, data)
}

// SetLateQuorumVoteExtension is a paid mutator transaction binding the contract method 0x330df7ff.
//
// Solidity: function setLateQuorumVoteExtension(uint48 newVoteExtension) returns()
func (_GlifGovernor *GlifGovernorTransactor) SetLateQuorumVoteExtension(opts *bind.TransactOpts, newVoteExtension *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "setLateQuorumVoteExtension", newVoteExtension)
}

// SetLateQuorumVoteExtension is a paid mutator transaction binding the contract method 0x330df7ff.
//
// Solidity: function setLateQuorumVoteExtension(uint48 newVoteExtension) returns()
func (_GlifGovernor *GlifGovernorSession) SetLateQuorumVoteExtension(newVoteExtension *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.Contract.SetLateQuorumVoteExtension(&_GlifGovernor.TransactOpts, newVoteExtension)
}

// SetLateQuorumVoteExtension is a paid mutator transaction binding the contract method 0x330df7ff.
//
// Solidity: function setLateQuorumVoteExtension(uint48 newVoteExtension) returns()
func (_GlifGovernor *GlifGovernorTransactorSession) SetLateQuorumVoteExtension(newVoteExtension *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.Contract.SetLateQuorumVoteExtension(&_GlifGovernor.TransactOpts, newVoteExtension)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_GlifGovernor *GlifGovernorTransactor) SetProposalThreshold(opts *bind.TransactOpts, newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "setProposalThreshold", newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_GlifGovernor *GlifGovernorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.Contract.SetProposalThreshold(&_GlifGovernor.TransactOpts, newProposalThreshold)
}

// SetProposalThreshold is a paid mutator transaction binding the contract method 0xece40cc1.
//
// Solidity: function setProposalThreshold(uint256 newProposalThreshold) returns()
func (_GlifGovernor *GlifGovernorTransactorSession) SetProposalThreshold(newProposalThreshold *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.Contract.SetProposalThreshold(&_GlifGovernor.TransactOpts, newProposalThreshold)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_GlifGovernor *GlifGovernorTransactor) SetVotingDelay(opts *bind.TransactOpts, newVotingDelay *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "setVotingDelay", newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_GlifGovernor *GlifGovernorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.Contract.SetVotingDelay(&_GlifGovernor.TransactOpts, newVotingDelay)
}

// SetVotingDelay is a paid mutator transaction binding the contract method 0x79051887.
//
// Solidity: function setVotingDelay(uint48 newVotingDelay) returns()
func (_GlifGovernor *GlifGovernorTransactorSession) SetVotingDelay(newVotingDelay *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.Contract.SetVotingDelay(&_GlifGovernor.TransactOpts, newVotingDelay)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_GlifGovernor *GlifGovernorTransactor) SetVotingPeriod(opts *bind.TransactOpts, newVotingPeriod uint32) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "setVotingPeriod", newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_GlifGovernor *GlifGovernorSession) SetVotingPeriod(newVotingPeriod uint32) (*types.Transaction, error) {
	return _GlifGovernor.Contract.SetVotingPeriod(&_GlifGovernor.TransactOpts, newVotingPeriod)
}

// SetVotingPeriod is a paid mutator transaction binding the contract method 0xe540d01d.
//
// Solidity: function setVotingPeriod(uint32 newVotingPeriod) returns()
func (_GlifGovernor *GlifGovernorTransactorSession) SetVotingPeriod(newVotingPeriod uint32) (*types.Transaction, error) {
	return _GlifGovernor.Contract.SetVotingPeriod(&_GlifGovernor.TransactOpts, newVotingPeriod)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_GlifGovernor *GlifGovernorTransactor) UpdateQuorumNumerator(opts *bind.TransactOpts, newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "updateQuorumNumerator", newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_GlifGovernor *GlifGovernorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.Contract.UpdateQuorumNumerator(&_GlifGovernor.TransactOpts, newQuorumNumerator)
}

// UpdateQuorumNumerator is a paid mutator transaction binding the contract method 0x06f3f9e6.
//
// Solidity: function updateQuorumNumerator(uint256 newQuorumNumerator) returns()
func (_GlifGovernor *GlifGovernorTransactorSession) UpdateQuorumNumerator(newQuorumNumerator *big.Int) (*types.Transaction, error) {
	return _GlifGovernor.Contract.UpdateQuorumNumerator(&_GlifGovernor.TransactOpts, newQuorumNumerator)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_GlifGovernor *GlifGovernorTransactor) UpdateTimelock(opts *bind.TransactOpts, newTimelock common.Address) (*types.Transaction, error) {
	return _GlifGovernor.contract.Transact(opts, "updateTimelock", newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_GlifGovernor *GlifGovernorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _GlifGovernor.Contract.UpdateTimelock(&_GlifGovernor.TransactOpts, newTimelock)
}

// UpdateTimelock is a paid mutator transaction binding the contract method 0xa890c910.
//
// Solidity: function updateTimelock(address newTimelock) returns()
func (_GlifGovernor *GlifGovernorTransactorSession) UpdateTimelock(newTimelock common.Address) (*types.Transaction, error) {
	return _GlifGovernor.Contract.UpdateTimelock(&_GlifGovernor.TransactOpts, newTimelock)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GlifGovernor *GlifGovernorTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GlifGovernor.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GlifGovernor *GlifGovernorSession) Receive() (*types.Transaction, error) {
	return _GlifGovernor.Contract.Receive(&_GlifGovernor.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_GlifGovernor *GlifGovernorTransactorSession) Receive() (*types.Transaction, error) {
	return _GlifGovernor.Contract.Receive(&_GlifGovernor.TransactOpts)
}

// GlifGovernorEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the GlifGovernor contract.
type GlifGovernorEIP712DomainChangedIterator struct {
	Event *GlifGovernorEIP712DomainChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorEIP712DomainChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorEIP712DomainChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorEIP712DomainChanged represents a EIP712DomainChanged event raised by the GlifGovernor contract.
type GlifGovernorEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GlifGovernor *GlifGovernorFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*GlifGovernorEIP712DomainChangedIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorEIP712DomainChangedIterator{contract: _GlifGovernor.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GlifGovernor *GlifGovernorFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *GlifGovernorEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorEIP712DomainChanged)
				if err := _GlifGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_GlifGovernor *GlifGovernorFilterer) ParseEIP712DomainChanged(log types.Log) (*GlifGovernorEIP712DomainChanged, error) {
	event := new(GlifGovernorEIP712DomainChanged)
	if err := _GlifGovernor.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorLateQuorumVoteExtensionSetIterator is returned from FilterLateQuorumVoteExtensionSet and is used to iterate over the raw logs and unpacked data for LateQuorumVoteExtensionSet events raised by the GlifGovernor contract.
type GlifGovernorLateQuorumVoteExtensionSetIterator struct {
	Event *GlifGovernorLateQuorumVoteExtensionSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorLateQuorumVoteExtensionSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorLateQuorumVoteExtensionSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorLateQuorumVoteExtensionSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorLateQuorumVoteExtensionSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorLateQuorumVoteExtensionSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorLateQuorumVoteExtensionSet represents a LateQuorumVoteExtensionSet event raised by the GlifGovernor contract.
type GlifGovernorLateQuorumVoteExtensionSet struct {
	OldVoteExtension uint64
	NewVoteExtension uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLateQuorumVoteExtensionSet is a free log retrieval operation binding the contract event 0x7ca4ac117ed3cdce75c1161d8207c440389b1a15d69d096831664657c07dafc2.
//
// Solidity: event LateQuorumVoteExtensionSet(uint64 oldVoteExtension, uint64 newVoteExtension)
func (_GlifGovernor *GlifGovernorFilterer) FilterLateQuorumVoteExtensionSet(opts *bind.FilterOpts) (*GlifGovernorLateQuorumVoteExtensionSetIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "LateQuorumVoteExtensionSet")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorLateQuorumVoteExtensionSetIterator{contract: _GlifGovernor.contract, event: "LateQuorumVoteExtensionSet", logs: logs, sub: sub}, nil
}

// WatchLateQuorumVoteExtensionSet is a free log subscription operation binding the contract event 0x7ca4ac117ed3cdce75c1161d8207c440389b1a15d69d096831664657c07dafc2.
//
// Solidity: event LateQuorumVoteExtensionSet(uint64 oldVoteExtension, uint64 newVoteExtension)
func (_GlifGovernor *GlifGovernorFilterer) WatchLateQuorumVoteExtensionSet(opts *bind.WatchOpts, sink chan<- *GlifGovernorLateQuorumVoteExtensionSet) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "LateQuorumVoteExtensionSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorLateQuorumVoteExtensionSet)
				if err := _GlifGovernor.contract.UnpackLog(event, "LateQuorumVoteExtensionSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseLateQuorumVoteExtensionSet is a log parse operation binding the contract event 0x7ca4ac117ed3cdce75c1161d8207c440389b1a15d69d096831664657c07dafc2.
//
// Solidity: event LateQuorumVoteExtensionSet(uint64 oldVoteExtension, uint64 newVoteExtension)
func (_GlifGovernor *GlifGovernorFilterer) ParseLateQuorumVoteExtensionSet(log types.Log) (*GlifGovernorLateQuorumVoteExtensionSet, error) {
	event := new(GlifGovernorLateQuorumVoteExtensionSet)
	if err := _GlifGovernor.contract.UnpackLog(event, "LateQuorumVoteExtensionSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorProposalCanceledIterator is returned from FilterProposalCanceled and is used to iterate over the raw logs and unpacked data for ProposalCanceled events raised by the GlifGovernor contract.
type GlifGovernorProposalCanceledIterator struct {
	Event *GlifGovernorProposalCanceled // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorProposalCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorProposalCanceled)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorProposalCanceled)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorProposalCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorProposalCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorProposalCanceled represents a ProposalCanceled event raised by the GlifGovernor contract.
type GlifGovernorProposalCanceled struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalCanceled is a free log retrieval operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_GlifGovernor *GlifGovernorFilterer) FilterProposalCanceled(opts *bind.FilterOpts) (*GlifGovernorProposalCanceledIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorProposalCanceledIterator{contract: _GlifGovernor.contract, event: "ProposalCanceled", logs: logs, sub: sub}, nil
}

// WatchProposalCanceled is a free log subscription operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_GlifGovernor *GlifGovernorFilterer) WatchProposalCanceled(opts *bind.WatchOpts, sink chan<- *GlifGovernorProposalCanceled) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "ProposalCanceled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorProposalCanceled)
				if err := _GlifGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCanceled is a log parse operation binding the contract event 0x789cf55be980739dad1d0699b93b58e806b51c9d96619bfa8fe0a28abaa7b30c.
//
// Solidity: event ProposalCanceled(uint256 proposalId)
func (_GlifGovernor *GlifGovernorFilterer) ParseProposalCanceled(log types.Log) (*GlifGovernorProposalCanceled, error) {
	event := new(GlifGovernorProposalCanceled)
	if err := _GlifGovernor.contract.UnpackLog(event, "ProposalCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorProposalCreatedIterator is returned from FilterProposalCreated and is used to iterate over the raw logs and unpacked data for ProposalCreated events raised by the GlifGovernor contract.
type GlifGovernorProposalCreatedIterator struct {
	Event *GlifGovernorProposalCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorProposalCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorProposalCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorProposalCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorProposalCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorProposalCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorProposalCreated represents a ProposalCreated event raised by the GlifGovernor contract.
type GlifGovernorProposalCreated struct {
	ProposalId  *big.Int
	Proposer    common.Address
	Targets     []common.Address
	Values      []*big.Int
	Signatures  []string
	Calldatas   [][]byte
	VoteStart   *big.Int
	VoteEnd     *big.Int
	Description string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProposalCreated is a free log retrieval operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_GlifGovernor *GlifGovernorFilterer) FilterProposalCreated(opts *bind.FilterOpts) (*GlifGovernorProposalCreatedIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorProposalCreatedIterator{contract: _GlifGovernor.contract, event: "ProposalCreated", logs: logs, sub: sub}, nil
}

// WatchProposalCreated is a free log subscription operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_GlifGovernor *GlifGovernorFilterer) WatchProposalCreated(opts *bind.WatchOpts, sink chan<- *GlifGovernorProposalCreated) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "ProposalCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorProposalCreated)
				if err := _GlifGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalCreated is a log parse operation binding the contract event 0x7d84a6263ae0d98d3329bd7b46bb4e8d6f98cd35a7adb45c274c8b7fd5ebd5e0.
//
// Solidity: event ProposalCreated(uint256 proposalId, address proposer, address[] targets, uint256[] values, string[] signatures, bytes[] calldatas, uint256 voteStart, uint256 voteEnd, string description)
func (_GlifGovernor *GlifGovernorFilterer) ParseProposalCreated(log types.Log) (*GlifGovernorProposalCreated, error) {
	event := new(GlifGovernorProposalCreated)
	if err := _GlifGovernor.contract.UnpackLog(event, "ProposalCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the GlifGovernor contract.
type GlifGovernorProposalExecutedIterator struct {
	Event *GlifGovernorProposalExecuted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorProposalExecuted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorProposalExecuted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorProposalExecuted represents a ProposalExecuted event raised by the GlifGovernor contract.
type GlifGovernorProposalExecuted struct {
	ProposalId *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_GlifGovernor *GlifGovernorFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*GlifGovernorProposalExecutedIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorProposalExecutedIterator{contract: _GlifGovernor.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_GlifGovernor *GlifGovernorFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *GlifGovernorProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorProposalExecuted)
				if err := _GlifGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 proposalId)
func (_GlifGovernor *GlifGovernorFilterer) ParseProposalExecuted(log types.Log) (*GlifGovernorProposalExecuted, error) {
	event := new(GlifGovernorProposalExecuted)
	if err := _GlifGovernor.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorProposalExtendedIterator is returned from FilterProposalExtended and is used to iterate over the raw logs and unpacked data for ProposalExtended events raised by the GlifGovernor contract.
type GlifGovernorProposalExtendedIterator struct {
	Event *GlifGovernorProposalExtended // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorProposalExtendedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorProposalExtended)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorProposalExtended)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorProposalExtendedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorProposalExtendedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorProposalExtended represents a ProposalExtended event raised by the GlifGovernor contract.
type GlifGovernorProposalExtended struct {
	ProposalId       *big.Int
	ExtendedDeadline uint64
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterProposalExtended is a free log retrieval operation binding the contract event 0x541f725fb9f7c98a30cc9c0ff32fbb14358cd7159c847a3aa20a2bdc442ba511.
//
// Solidity: event ProposalExtended(uint256 indexed proposalId, uint64 extendedDeadline)
func (_GlifGovernor *GlifGovernorFilterer) FilterProposalExtended(opts *bind.FilterOpts, proposalId []*big.Int) (*GlifGovernorProposalExtendedIterator, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "ProposalExtended", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return &GlifGovernorProposalExtendedIterator{contract: _GlifGovernor.contract, event: "ProposalExtended", logs: logs, sub: sub}, nil
}

// WatchProposalExtended is a free log subscription operation binding the contract event 0x541f725fb9f7c98a30cc9c0ff32fbb14358cd7159c847a3aa20a2bdc442ba511.
//
// Solidity: event ProposalExtended(uint256 indexed proposalId, uint64 extendedDeadline)
func (_GlifGovernor *GlifGovernorFilterer) WatchProposalExtended(opts *bind.WatchOpts, sink chan<- *GlifGovernorProposalExtended, proposalId []*big.Int) (event.Subscription, error) {

	var proposalIdRule []interface{}
	for _, proposalIdItem := range proposalId {
		proposalIdRule = append(proposalIdRule, proposalIdItem)
	}

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "ProposalExtended", proposalIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorProposalExtended)
				if err := _GlifGovernor.contract.UnpackLog(event, "ProposalExtended", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalExtended is a log parse operation binding the contract event 0x541f725fb9f7c98a30cc9c0ff32fbb14358cd7159c847a3aa20a2bdc442ba511.
//
// Solidity: event ProposalExtended(uint256 indexed proposalId, uint64 extendedDeadline)
func (_GlifGovernor *GlifGovernorFilterer) ParseProposalExtended(log types.Log) (*GlifGovernorProposalExtended, error) {
	event := new(GlifGovernorProposalExtended)
	if err := _GlifGovernor.contract.UnpackLog(event, "ProposalExtended", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorProposalQueuedIterator is returned from FilterProposalQueued and is used to iterate over the raw logs and unpacked data for ProposalQueued events raised by the GlifGovernor contract.
type GlifGovernorProposalQueuedIterator struct {
	Event *GlifGovernorProposalQueued // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorProposalQueuedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorProposalQueued)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorProposalQueued)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorProposalQueuedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorProposalQueuedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorProposalQueued represents a ProposalQueued event raised by the GlifGovernor contract.
type GlifGovernorProposalQueued struct {
	ProposalId *big.Int
	EtaSeconds *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProposalQueued is a free log retrieval operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_GlifGovernor *GlifGovernorFilterer) FilterProposalQueued(opts *bind.FilterOpts) (*GlifGovernorProposalQueuedIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorProposalQueuedIterator{contract: _GlifGovernor.contract, event: "ProposalQueued", logs: logs, sub: sub}, nil
}

// WatchProposalQueued is a free log subscription operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_GlifGovernor *GlifGovernorFilterer) WatchProposalQueued(opts *bind.WatchOpts, sink chan<- *GlifGovernorProposalQueued) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "ProposalQueued")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorProposalQueued)
				if err := _GlifGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalQueued is a log parse operation binding the contract event 0x9a2e42fd6722813d69113e7d0079d3d940171428df7373df9c7f7617cfda2892.
//
// Solidity: event ProposalQueued(uint256 proposalId, uint256 etaSeconds)
func (_GlifGovernor *GlifGovernorFilterer) ParseProposalQueued(log types.Log) (*GlifGovernorProposalQueued, error) {
	event := new(GlifGovernorProposalQueued)
	if err := _GlifGovernor.contract.UnpackLog(event, "ProposalQueued", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorProposalThresholdSetIterator is returned from FilterProposalThresholdSet and is used to iterate over the raw logs and unpacked data for ProposalThresholdSet events raised by the GlifGovernor contract.
type GlifGovernorProposalThresholdSetIterator struct {
	Event *GlifGovernorProposalThresholdSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorProposalThresholdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorProposalThresholdSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorProposalThresholdSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorProposalThresholdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorProposalThresholdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorProposalThresholdSet represents a ProposalThresholdSet event raised by the GlifGovernor contract.
type GlifGovernorProposalThresholdSet struct {
	OldProposalThreshold *big.Int
	NewProposalThreshold *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterProposalThresholdSet is a free log retrieval operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_GlifGovernor *GlifGovernorFilterer) FilterProposalThresholdSet(opts *bind.FilterOpts) (*GlifGovernorProposalThresholdSetIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorProposalThresholdSetIterator{contract: _GlifGovernor.contract, event: "ProposalThresholdSet", logs: logs, sub: sub}, nil
}

// WatchProposalThresholdSet is a free log subscription operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_GlifGovernor *GlifGovernorFilterer) WatchProposalThresholdSet(opts *bind.WatchOpts, sink chan<- *GlifGovernorProposalThresholdSet) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "ProposalThresholdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorProposalThresholdSet)
				if err := _GlifGovernor.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProposalThresholdSet is a log parse operation binding the contract event 0xccb45da8d5717e6c4544694297c4ba5cf151d455c9bb0ed4fc7a38411bc05461.
//
// Solidity: event ProposalThresholdSet(uint256 oldProposalThreshold, uint256 newProposalThreshold)
func (_GlifGovernor *GlifGovernorFilterer) ParseProposalThresholdSet(log types.Log) (*GlifGovernorProposalThresholdSet, error) {
	event := new(GlifGovernorProposalThresholdSet)
	if err := _GlifGovernor.contract.UnpackLog(event, "ProposalThresholdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorQuorumNumeratorUpdatedIterator is returned from FilterQuorumNumeratorUpdated and is used to iterate over the raw logs and unpacked data for QuorumNumeratorUpdated events raised by the GlifGovernor contract.
type GlifGovernorQuorumNumeratorUpdatedIterator struct {
	Event *GlifGovernorQuorumNumeratorUpdated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorQuorumNumeratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorQuorumNumeratorUpdated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorQuorumNumeratorUpdated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorQuorumNumeratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorQuorumNumeratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorQuorumNumeratorUpdated represents a QuorumNumeratorUpdated event raised by the GlifGovernor contract.
type GlifGovernorQuorumNumeratorUpdated struct {
	OldQuorumNumerator *big.Int
	NewQuorumNumerator *big.Int
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterQuorumNumeratorUpdated is a free log retrieval operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_GlifGovernor *GlifGovernorFilterer) FilterQuorumNumeratorUpdated(opts *bind.FilterOpts) (*GlifGovernorQuorumNumeratorUpdatedIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorQuorumNumeratorUpdatedIterator{contract: _GlifGovernor.contract, event: "QuorumNumeratorUpdated", logs: logs, sub: sub}, nil
}

// WatchQuorumNumeratorUpdated is a free log subscription operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_GlifGovernor *GlifGovernorFilterer) WatchQuorumNumeratorUpdated(opts *bind.WatchOpts, sink chan<- *GlifGovernorQuorumNumeratorUpdated) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "QuorumNumeratorUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorQuorumNumeratorUpdated)
				if err := _GlifGovernor.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseQuorumNumeratorUpdated is a log parse operation binding the contract event 0x0553476bf02ef2726e8ce5ced78d63e26e602e4a2257b1f559418e24b4633997.
//
// Solidity: event QuorumNumeratorUpdated(uint256 oldQuorumNumerator, uint256 newQuorumNumerator)
func (_GlifGovernor *GlifGovernorFilterer) ParseQuorumNumeratorUpdated(log types.Log) (*GlifGovernorQuorumNumeratorUpdated, error) {
	event := new(GlifGovernorQuorumNumeratorUpdated)
	if err := _GlifGovernor.contract.UnpackLog(event, "QuorumNumeratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorTimelockChangeIterator is returned from FilterTimelockChange and is used to iterate over the raw logs and unpacked data for TimelockChange events raised by the GlifGovernor contract.
type GlifGovernorTimelockChangeIterator struct {
	Event *GlifGovernorTimelockChange // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorTimelockChangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorTimelockChange)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorTimelockChange)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorTimelockChangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorTimelockChangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorTimelockChange represents a TimelockChange event raised by the GlifGovernor contract.
type GlifGovernorTimelockChange struct {
	OldTimelock common.Address
	NewTimelock common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTimelockChange is a free log retrieval operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_GlifGovernor *GlifGovernorFilterer) FilterTimelockChange(opts *bind.FilterOpts) (*GlifGovernorTimelockChangeIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorTimelockChangeIterator{contract: _GlifGovernor.contract, event: "TimelockChange", logs: logs, sub: sub}, nil
}

// WatchTimelockChange is a free log subscription operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_GlifGovernor *GlifGovernorFilterer) WatchTimelockChange(opts *bind.WatchOpts, sink chan<- *GlifGovernorTimelockChange) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "TimelockChange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorTimelockChange)
				if err := _GlifGovernor.contract.UnpackLog(event, "TimelockChange", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTimelockChange is a log parse operation binding the contract event 0x08f74ea46ef7894f65eabfb5e6e695de773a000b47c529ab559178069b226401.
//
// Solidity: event TimelockChange(address oldTimelock, address newTimelock)
func (_GlifGovernor *GlifGovernorFilterer) ParseTimelockChange(log types.Log) (*GlifGovernorTimelockChange, error) {
	event := new(GlifGovernorTimelockChange)
	if err := _GlifGovernor.contract.UnpackLog(event, "TimelockChange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorVoteCastIterator is returned from FilterVoteCast and is used to iterate over the raw logs and unpacked data for VoteCast events raised by the GlifGovernor contract.
type GlifGovernorVoteCastIterator struct {
	Event *GlifGovernorVoteCast // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorVoteCastIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorVoteCast)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorVoteCast)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorVoteCastIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorVoteCastIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorVoteCast represents a VoteCast event raised by the GlifGovernor contract.
type GlifGovernorVoteCast struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCast is a free log retrieval operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_GlifGovernor *GlifGovernorFilterer) FilterVoteCast(opts *bind.FilterOpts, voter []common.Address) (*GlifGovernorVoteCastIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return &GlifGovernorVoteCastIterator{contract: _GlifGovernor.contract, event: "VoteCast", logs: logs, sub: sub}, nil
}

// WatchVoteCast is a free log subscription operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_GlifGovernor *GlifGovernorFilterer) WatchVoteCast(opts *bind.WatchOpts, sink chan<- *GlifGovernorVoteCast, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "VoteCast", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorVoteCast)
				if err := _GlifGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCast is a log parse operation binding the contract event 0xb8e138887d0aa13bab447e82de9d5c1777041ecd21ca36ba824ff1e6c07ddda4.
//
// Solidity: event VoteCast(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason)
func (_GlifGovernor *GlifGovernorFilterer) ParseVoteCast(log types.Log) (*GlifGovernorVoteCast, error) {
	event := new(GlifGovernorVoteCast)
	if err := _GlifGovernor.contract.UnpackLog(event, "VoteCast", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorVoteCastWithParamsIterator is returned from FilterVoteCastWithParams and is used to iterate over the raw logs and unpacked data for VoteCastWithParams events raised by the GlifGovernor contract.
type GlifGovernorVoteCastWithParamsIterator struct {
	Event *GlifGovernorVoteCastWithParams // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorVoteCastWithParamsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorVoteCastWithParams)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorVoteCastWithParams)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorVoteCastWithParamsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorVoteCastWithParamsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorVoteCastWithParams represents a VoteCastWithParams event raised by the GlifGovernor contract.
type GlifGovernorVoteCastWithParams struct {
	Voter      common.Address
	ProposalId *big.Int
	Support    uint8
	Weight     *big.Int
	Reason     string
	Params     []byte
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterVoteCastWithParams is a free log retrieval operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_GlifGovernor *GlifGovernorFilterer) FilterVoteCastWithParams(opts *bind.FilterOpts, voter []common.Address) (*GlifGovernorVoteCastWithParamsIterator, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return &GlifGovernorVoteCastWithParamsIterator{contract: _GlifGovernor.contract, event: "VoteCastWithParams", logs: logs, sub: sub}, nil
}

// WatchVoteCastWithParams is a free log subscription operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_GlifGovernor *GlifGovernorFilterer) WatchVoteCastWithParams(opts *bind.WatchOpts, sink chan<- *GlifGovernorVoteCastWithParams, voter []common.Address) (event.Subscription, error) {

	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "VoteCastWithParams", voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorVoteCastWithParams)
				if err := _GlifGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVoteCastWithParams is a log parse operation binding the contract event 0xe2babfbac5889a709b63bb7f598b324e08bc5a4fb9ec647fb3cbc9ec07eb8712.
//
// Solidity: event VoteCastWithParams(address indexed voter, uint256 proposalId, uint8 support, uint256 weight, string reason, bytes params)
func (_GlifGovernor *GlifGovernorFilterer) ParseVoteCastWithParams(log types.Log) (*GlifGovernorVoteCastWithParams, error) {
	event := new(GlifGovernorVoteCastWithParams)
	if err := _GlifGovernor.contract.UnpackLog(event, "VoteCastWithParams", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorVotingDelaySetIterator is returned from FilterVotingDelaySet and is used to iterate over the raw logs and unpacked data for VotingDelaySet events raised by the GlifGovernor contract.
type GlifGovernorVotingDelaySetIterator struct {
	Event *GlifGovernorVotingDelaySet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorVotingDelaySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorVotingDelaySet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorVotingDelaySet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorVotingDelaySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorVotingDelaySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorVotingDelaySet represents a VotingDelaySet event raised by the GlifGovernor contract.
type GlifGovernorVotingDelaySet struct {
	OldVotingDelay *big.Int
	NewVotingDelay *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterVotingDelaySet is a free log retrieval operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_GlifGovernor *GlifGovernorFilterer) FilterVotingDelaySet(opts *bind.FilterOpts) (*GlifGovernorVotingDelaySetIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorVotingDelaySetIterator{contract: _GlifGovernor.contract, event: "VotingDelaySet", logs: logs, sub: sub}, nil
}

// WatchVotingDelaySet is a free log subscription operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_GlifGovernor *GlifGovernorFilterer) WatchVotingDelaySet(opts *bind.WatchOpts, sink chan<- *GlifGovernorVotingDelaySet) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "VotingDelaySet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorVotingDelaySet)
				if err := _GlifGovernor.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVotingDelaySet is a log parse operation binding the contract event 0xc565b045403dc03c2eea82b81a0465edad9e2e7fc4d97e11421c209da93d7a93.
//
// Solidity: event VotingDelaySet(uint256 oldVotingDelay, uint256 newVotingDelay)
func (_GlifGovernor *GlifGovernorFilterer) ParseVotingDelaySet(log types.Log) (*GlifGovernorVotingDelaySet, error) {
	event := new(GlifGovernorVotingDelaySet)
	if err := _GlifGovernor.contract.UnpackLog(event, "VotingDelaySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// GlifGovernorVotingPeriodSetIterator is returned from FilterVotingPeriodSet and is used to iterate over the raw logs and unpacked data for VotingPeriodSet events raised by the GlifGovernor contract.
type GlifGovernorVotingPeriodSetIterator struct {
	Event *GlifGovernorVotingPeriodSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *GlifGovernorVotingPeriodSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GlifGovernorVotingPeriodSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(GlifGovernorVotingPeriodSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *GlifGovernorVotingPeriodSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GlifGovernorVotingPeriodSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GlifGovernorVotingPeriodSet represents a VotingPeriodSet event raised by the GlifGovernor contract.
type GlifGovernorVotingPeriodSet struct {
	OldVotingPeriod *big.Int
	NewVotingPeriod *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterVotingPeriodSet is a free log retrieval operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_GlifGovernor *GlifGovernorFilterer) FilterVotingPeriodSet(opts *bind.FilterOpts) (*GlifGovernorVotingPeriodSetIterator, error) {

	logs, sub, err := _GlifGovernor.contract.FilterLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return &GlifGovernorVotingPeriodSetIterator{contract: _GlifGovernor.contract, event: "VotingPeriodSet", logs: logs, sub: sub}, nil
}

// WatchVotingPeriodSet is a free log subscription operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_GlifGovernor *GlifGovernorFilterer) WatchVotingPeriodSet(opts *bind.WatchOpts, sink chan<- *GlifGovernorVotingPeriodSet) (event.Subscription, error) {

	logs, sub, err := _GlifGovernor.contract.WatchLogs(opts, "VotingPeriodSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GlifGovernorVotingPeriodSet)
				if err := _GlifGovernor.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseVotingPeriodSet is a log parse operation binding the contract event 0x7e3f7f0708a84de9203036abaa450dccc85ad5ff52f78c170f3edb55cf5e8828.
//
// Solidity: event VotingPeriodSet(uint256 oldVotingPeriod, uint256 newVotingPeriod)
func (_GlifGovernor *GlifGovernorFilterer) ParseVotingPeriodSet(log types.Log) (*GlifGovernorVotingPeriodSet, error) {
	event := new(GlifGovernorVotingPeriodSet)
	if err := _GlifGovernor.contract.UnpackLog(event, "VotingPeriodSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
