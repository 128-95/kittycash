package visor

import "github.com/skycoin/skycoin/src/coin"

const (
	// Maximum supply of skycoins
	MaxCoinSupply uint64 = 1e8 // 100,000,000 million

	// Number of distribution addresses
	DistributionAddressesTotal uint64 = 100

	DistributionAddressInitialBalance uint64 = MaxCoinSupply / DistributionAddressesTotal

	// Initial number of unlocked addresses
	InitialUnlockedCount uint64 = 100

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
	"2DehfGNYtk9AFFNELQ3jk7GyFrtJC4RJf87",
	"nyiuJiVtDm1WTRVUSFpuEMeXnpQ1MvqEtF",
	"2WqM7oEFfVPND5gzK168ty8vuBEejQymDSV",
	"79GjMYDjRNG3TLzjxBe513EfKQgis7FZbj",
	"2WJWN4pcBmGuXfBPjbELiDt2wkL64P6LJoQ",
	"5DZFE6dvSHSV5zjWqqR2cvStU2n7zJTDMr",
	"2iQYNigbiyidQT1487rnwihrUsy3xpmrjah",
	"2azMgMgvEqn9gpnkUTq596J7PKpQP7JD1H5",
	"2ZKhM9nrcQeJNt2DuhkJZGcMKgdzALeSTYz",
	"4jnmRr9SXThjYHTex77sSx1EGjX2CNrvEB",
	"9V3fLPvB1dBnNRPPpLDA14bbsQt7363byQ",
	"qAvSjyJvZDa8oB6htaW57ZPzszcZk7BDQE",
	"pbu1kWN9LZXNhCRtxsef7RhaSCeQA9fsCz",
	"2Ad3HME7gXowHX4CvPi71URiPDE6BFakB3d",
	"23SbdmjG5PRYYDfzSW2HyALUD3AzRZ5KxdP",
	"gju6y5PvSaut3aoo518d62BR6udGMmZ7uk",
	"YvNi459Xeg4nuem1qLChKAU431QvQxh6QP",
	"4Jkhu4zTWzB41zXW4ZzQChgZJxWk8RutpJ",
	"28Zou6bEwZD8MFrMFkEot9v8KVcr3eigY6V",
	"2kPW47GenVEsM2wpEda963viduACxZYSFiA",
	"tjTPdeuNX8SW9syTt5W3Z3ndGkPjbLfwK",
	"222MvhS11Lu4Rj6Dr8gYgTZmDtKrCfvjuuH",
	"CMkiYXyMSzAAEY8KvLd1129NQ3xScKzFNz",
	"2Fg5y5yS14Tk3Kzcc19BAuKbkPLSo8F9Skd",
	"2NDtm5uMgJWP2g9sP7hXcifG9z26EJfWgKU",
	"6M1MhEc32dpGmor8Lr38F4bkUQPsn21Ua8",
	"hh86GQ5qSSh8LsHLsvsEdGAHqtxWVWZKa8",
	"epLj8mLu7LAXCb9kCHhendU8UAR82hoKrZ",
	"2SqgnBvw2CLwbpeWvauhWPh5sngQNgBteLH",
	"GcK46Cu2WQ8LsmMmcHAJfufmMdXrLev8jw",
	"Czs9WC1XV4gi4bRvbMsVHybVZAzHNHwB1p",
	"2GroEZt945FCyqa25JH87JHctzQMMNgQYEC",
	"cb4nNVXLMxK3F7aVP1M2BUM7C1DZEjpGvy",
	"CKcdzNHUtjKPUABQH5thV9gV918x6ZLxTy",
	"25mQXfTcsghVPi6rkn6D8akSKsnrZVCAszD",
	"Em2s7acJsS3PxHmHJ8MCmx1tc9HFkxVvxV",
	"23DkSg5MAoreFF2RW7oExjbgLspWReDevd6",
	"Z1J9oQ1anzSe9DTLiNVBKYYqieHdW3DNkz",
	"2NQipsWVcGykFHBFwofBL8FgkybqvrUPJVb",
	"VW69UpUrJijiqt6MFUNcoDiMPAYTeweEqE",
	"oPnG1ZCsWkfxhp2Gjvt9njNpyGCT3YpqdU",
	"z5wZBYTRHrf8ybTeUeftEqMQkN8WuBgUKC",
	"2eT5yVsZUBKFwswW7ckePGNiFtgwG9cTtpv",
	"WEUFkx7R5yy8QhqVCTfkkobLqdgW3g2Rmf",
	"nZS47bMJ4vifeHnD1NvnSEH9JjvYWWofQF",
	"2WqPLHQ7s1RtKNzbtr3PfWxuzoVHsWzgR6m",
	"AakeCYqyadmqXvgmbCwJmbPpA17seXCQFD",
	"26TWjjPTvqX9GgM7yKPei2KvxrXfzbDJsNR",
	"qKYn4hzXZwKQb4KZJe2CmPHXNb4MjWmPSc",
	"bfuGeJMeSRCxRyTobHwCdv8Q9gMHoXmBhF",
	"2dQaQVgqXgkj6aarjSCwoVKNuh3Yizygpq1",
	"2BDspMwyT9vf92fE2eqrx8XhmbxEofxfrnE",
	"252MKmpjjuDRwnqKmJBYf7wwEjGSWLqHpFo",
	"pS4JytexuPawXtdPku3fZF8AZCy5UhQ6zc",
	"2gdPirEL75zpWwdP2sETNjw6omedTRCAYf9",
	"2QPXJ8MCvLeUeBp1po5zmv886QiusuFrWDi",
	"KoZNq97AjcTcAGmwwoLBKu5vpzRCFnZzhE",
	"2XWgLnWnA57iNEPEc224pJpCdXEhgyifYLj",
	"Xp7pVydwwALUGYVvs2A1k3TTbwC3mKgS5x",
	"24NJrgp1digmzK3bffFnN7NmDmBWSsEPhiv",
	"iqgzmd79zMmF8PAgu3zkaUJ5iUnexwfHtL",
	"ZkhZjRrvXjdhbb6N3ZhV6yEbtLwjDcpyCE",
	"366RPcHKPoEwuQqGrxx2ThvwMHLvmhCxgY",
	"2ca5WDkoaJu6uqyTXMBMXP2LVeDaqVHUjsJ",
	"2SjS8xVt28R6RZ13suMedvTrKCsUYbH7qX8",
	"2iZYioQ99dys9CWamYSTtjeLTTtAWube3yA",
	"2RPzDy3ncdBW9uZMxqMMDBAUqZiPmfVCGUd",
	"2VBEgqamynB4xYtHSHJDskpxJLi3D3fnQNu",
	"2eR5eUYmR3KJ4Qypd3BvMbUgam8mwBzpoEJ",
	"5HdJBz2aWT7Bd9yzBrYhKdE8mPjp6UpKsd",
	"h9Ue9mnNjrZ1cTFBbtDeFXmceZgreFfJpS",
	"cxAxAu7RyK7QGPg138a2idDVFVBRoWdvWi",
	"ca8HV4iCF1kqnNsefK6FT8LQswNHREmMUt",
	"P2GqMv7bMHFEqR4FvEpgW2bbSQHF8LY3kL",
	"2YhjBsd9LqfZz1rs15gLkp7YbNMekcD3wAK",
	"FcgPxBQvZd5UMk1RaQSqHqhKSPKrsUShNh",
	"2hgWtQHQdjy6ajRe1N2GkjvVes28P7QL2Ah",
	"2M4BEvVQt2iJtEbtiJwp48E11xLin5SQ4k7",
	"Y5picGkLVLBHE2YLa7EjX2S2Qe2kqrTQPT",
	"oRSTfd8hHw7Vpr3aSJVGDwG4dya57ADqWd",
	"HkinEhD7Q9Mu7MUtX42wTnXEuwS1epyDq2",
	"27zUcMCjS8tHxFcTwGA3uWYMhA36Rt4vFjj",
	"2Vnp4hGfc4xE99uNHunnicihHEymbf4cTHf",
	"2iyXu8kCYcktLrHhtw973uERe53JbtHKP1h",
	"2RdaacNoyVWqno32vJ1S2oEPgbKPbK7LAUc",
	"28Az6ka4LfZpQBiqHZBxGTDDKoEBvecfD3B",
	"TnFP1Q9cERfNnpoocygEQjHLiHVfXrA2Vy",
	"27SPnC6WkhBx33iBM8ms1oEkHs6bet4f8cz",
	"CG1KGYencqKWv9LcoutxL3tf5gysBpy94T",
	"2bXgFVCk5hmesa3hfvZFkz62fsCg7b6wJ2d",
	"Maht1n2f2MrbnrqikNLbcJQDVdC9ByfxbD",
	"AJLkgsyW4ZC7FrdsjPYzgq5sKKk25mcnjr",
	"Ye9C1KnwdfnBT5kVjgg1PFpmcuSs6FF1eK",
	"S5vALMa9GCGmQoAJkt7MQ5jkoaLwKnP1rr",
	"29tTZYbE2bCBMD1CiR5bwitXnsKmLjdXgzV",
	"MV8faKwtNBZQ5iLezgvYPnheDpXbCBim5i",
	"2j2GUziP24kUhvj61rX489ZtW2MyUoVPk5r",
	"oTHminHxBRT91KNRytdipxngYYDdzNxHT9",
	"tk3kwRBWzy8oUvmg7jD2KQejSksAZr7Et9",
	"2jUL7gXnMdKvxMpT6o4upg12KAYfMzT58Qw",
}
