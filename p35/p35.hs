-- Problem 35: Circular primes
--
-- https://projecteuler.net/problem=35
--
-- The number, 197, is called a circular prime because all rotations of the digits: 197, 971, and 719, are themselves prime.
--
-- There are thirteen such primes below 100: 2, 3, 5, 7, 11, 13, 17, 31, 37, 71, 73, 79, and 97.
--
-- How many circular primes are there below one million?

import Data.Char (digitToInt)

primesUnder n = primeSieve 2 0 n [2..(n - 1)]

primeSieve p i n xs
    | p * p > n = xs
    | p < n = primeSieve (xs !! i) (i + 1) n (filter (\x -> (x `mod` p /= 0) || (x == p)) xs)
    | otherwise = xs

circularCombos n = circularCombos' n (tail n ++ [head n]) []
circularCombos' i n cs
    | i /= n = circularCombos' i (tail n ++ [head n]) (cs ++ [(read n)])
    | otherwise = cs

circularPrime n ps = foldl (&&) True [elem c ps | c <- circularCombos (show n)]

circularPrimes ps = filter (\n -> circularPrime n ps) ps

main = putStrLn $ show $ length $ circularPrimes (primesUnder 1000000)
