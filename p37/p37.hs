-- Problem 37: Truncatable Primes
--
-- https://projecteuler.net/problem=37
--
-- The number 3797 has an interesting property. Being prime itself, it is possible to continuously remove digits from left to right, 
-- and remain prime at each stage: 3797, 797, 97, and 7. Similarly we can work from right to left: 3797, 379, 37, and 3.
--
-- Find the sum of the only eleven primes that are both truncatable from left to right and right to left.
--
-- NOTE: 2, 3, 5, and 7 are not considered to be truncatable primes.

import Data.List

primesUnder n = sieve 2 0 n [2..(n - 1)]

sieve p i n xs
    | p * p > n = xs
    | p < n = sieve (xs !! i) (i + 1) n (filter (\x -> (x `mod` p /= 0) || (x == p)) xs)
    | otherwise = xs

truncateStringLeft p tp = if length p == 1 then tp else truncateStringLeft (tail p) (tp ++ [tail p])
truncateLeft :: Integer -> [Integer]
truncateLeft p = map read $ truncateStringLeft (show p) []

truncateStringRight p tp = if length p == 1 then tp else truncateStringRight (reverse $ tail $ reverse p) (tp ++ [reverse $ tail $ reverse p])
truncateRight :: Integer -> [Integer]
truncateRight p = map read $ truncateStringRight (show p) []

truncatablePrime p ps = if all (\k -> elem k ps) (truncateLeft p ++ truncateRight p) then [p] else []

truncatablePrimes (p:ps) ts
    | ps == [] = filter (\k -> k > 10) $ nub ts
    | otherwise = truncatablePrimes ps (ts ++ truncatablePrime p ps)

main = putStrLn $ show $ sum $ truncatablePrimes (reverse $ sort $ primesUnder 800000) []
