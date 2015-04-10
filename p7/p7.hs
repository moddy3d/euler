-- Problem 7: 10001st prime
--
-- https://projecteuler.net/problem=7
--
-- By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
--
-- What is the 10001st prime number?

prime n = (take 2 $ filter (\x -> n `mod` x == 0) [1..n]) !! 1 == n

primesAt n = primes' n 1 2
primes' n i p
    | (prime p) && (i < n)  = primes' n (i + 1) (p + 1)
    | (prime p) && (i == n)  = p
    | otherwise = primes' n i (p + 1)

main = putStrLn $ show $ primesAt 10001
