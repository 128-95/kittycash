package visor

import "github.com/skycoin/skycoin/src/coin"

const (
	// Maximum supply of skycoins
	MaxCoinSupply uint64 = 21e6 // 21,000,000 million

	// Number of distribution addresses
	DistributionAddressesTotal uint64 = 21

	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal

	// Initial number of unlocked addresses
	InitialUnlockedCount uint64 = 21

	// Number of addresses to unlock per unlock time interval
	UnlockAddressRate uint64 = 5

	// Unlock time interval, measured in seconds
	// Once the InitialUnlockedCount is exhausted,
	// UnlockAddressRate addresses will be unlocked per UnlockTimeInterval
	UnlockTimeInterval uint64 = 60 * 60 * 24 * 365 // 1 year
)

func init() {
	if MaxCoinSupply%DistributionAddressesTotal != 0 {
		panic("MaxCoinSupply should be perfectly divisible by DistributionAddressesTotal")
	}
}

// Returns a copy of the hardcoded distribution addresses array.
// Each address has 1,000,000 coins. There are 100 addresses.
func GetDistributionAddresses() []string {
	addrs := make([]string, len(distributionAddresses))
	for i := range distributionAddresses {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are unlocked, i.e. they have spendable outputs
func GetUnlockedDistributionAddresses() []string {
	// The first InitialUnlockedCount (25) addresses are unlocked by default.
	// Subsequent addresses will be unlocked at a rate of UnlockAddressRate (5) per year,
	// after the InitialUnlockedCount (25) addresses have no remaining balance.
	// The unlock timer will be enabled manually once the
	// InitialUnlockedCount (25) addresses are distributed.

	// NOTE: To have automatic unlocking, transaction verification would have
	// to be handled in visor rather than in coin.Transactions.Visor(), because
	// the coin package is agnostic to the state of the blockchain and cannot reference it.
	// Instead of automatic unlocking, we can hardcode the timestamp at which the first 30%
	// is distributed, then compute the unlocked addresses easily here.

	addrs := make([]string, InitialUnlockedCount)
	for i := range distributionAddresses[:InitialUnlockedCount] {
		addrs[i] = distributionAddresses[i]
	}
	return addrs
}

// Returns distribution addresses that are locked, i.e. they have unspendable outputs
func GetLockedDistributionAddresses() []string {
	// TODO -- once we reach 30% distribution, we can hardcode the
	// initial timestamp for releasing more coins
	addrs := make([]string, DistributionAddressesTotal-InitialUnlockedCount)
	for i := range distributionAddresses[InitialUnlockedCount:] {
		addrs[i] = distributionAddresses[InitialUnlockedCount+uint64(i)]
	}
	return addrs
}

// Returns true if the transaction spends locked outputs
func TransactionIsLocked(inUxs coin.UxArray) bool {
	lockedAddrs := GetLockedDistributionAddresses()
	lockedAddrsMap := make(map[string]struct{})
	for _, a := range lockedAddrs {
		lockedAddrsMap[a] = struct{}{}
	}

	for _, o := range inUxs {
		uxAddr := o.Body.Address.String()
		if _, ok := lockedAddrsMap[uxAddr]; ok {
			return true
		}
	}

	return false
}

var distributionAddresses = [DistributionAddressesTotal]string{
	"6mqPS39siowV2ja5oPz4qRRBw7ZAc7wq2y",
	"n88zyADYmRDjy8qZmxy3KXyu6wDduUZzWd",
	"o88cBBjTb1xKVtvq5ddGLWcRcQPZ1b2Fec",
	"2Xb3akZLA3KzA7bvWg5hDEj2rqqREXnNRre",
	"2AD2nJdP3YpVrwtRxZHGWFpkp6jt6gfVSZb",
	"1Ge6g5ZjxMJQVxeF7ou8YPub2cCJ2BXVWa",
	"BDvfdxpCjZm3WbEL3aFWGZNbDUXVT9NYwm",
	"j8dcfLQxF3QT5Hff3cEfQR9jUCP8XKfEsa",
	"2cawVvbRLbCaJunNzT51aEWCybfGRDAaKAv",
	"23RVKJERrnMXUesKgCmhfVSkTRwELBTiN7A",
	"2TEXDNUJRJcxjKYh9usKJo7xPZF3ZV3NfsC",
	"XhWKHDpbSZwy31Y72GCQPxnXyRAwo4TD56",
	"ULx6swbsvApUztD3nTHdLCSKWoDHHhSUfR",
	"2LGZVTbeYuV9g7e4UxXuQE83ETZ2SYYVMem",
	"PEdcCgUGxWKxb2LtSQAbCsdbthSCYfomSF",
	"Bs2B2RFFYAWnGQBp45d54myc34eHqD476w",
	"2GdPf3ZMzWY7oF7mPVsXBMGrUhmN4QWe7eY",
	"QZP3mFz7yc1wc6Kg6zo1ioRr8WrUweEbHJ",
	"2R5YXYRE3SgzCX6gfQTwasLzjPvqp8QcoW1",
	"PRJRkJyc4i9J2ow8BAEbkXeQdtELxMVAov",
	"2CRnU7WH5qcakV1yEAf8rjgU3erEYidj8X1",
}
