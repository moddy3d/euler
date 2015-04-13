-- Problem 14: Longest Collatz Sequence
--
-- https://projecteuler.net/problem=14
--
-- The following iterative sequence is defined for the set of positive integers:
-- n → n/2 (n is even)
-- n → 3n + 1 (n is odd)
--
-- Using the rule above and starting with 13, we generate the following sequence:
-- 13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
--
-- It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms.
-- Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.
--
-- Which starting number, under one million, produces the longest chain?
--
-- NOTE: Once the chain starts the terms are allowed to go above one million.

collatz n = if n `mod` 2 == 0 then n `div` 2 else 3 * n + 1

collatzCount n = collatzCount' n 0
collatzCount' n c
    | collatz n == 1 = c
    | otherwise = collatzCount' (collatz n) (c + 1)

maximum' [] = error "maximum of empty list"
maximum' (x:xs) = maxTail x xs
    where maxTail currentMax [] = currentMax
          maxTail (m, n) (p:ps)
            | n < (snd p) = maxTail p ps
            | otherwise = maxTail (m, n) ps

main = putStrLn $ show $ maximum $ zip (map collatzCount [1..1000000]) [1..1000000] 
