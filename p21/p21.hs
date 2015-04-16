-- Problem 21: Amicable numbers
--
-- https://projecteuler.net/problem=21
-- 
-- Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
-- If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.
--
-- For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.
--
-- Evaluate the sum of all the amicable numbers under 10000.

import Data.List

factors n = lows ++ (reverse $ map (div n) lows)
    where lows = filter ((== 0) . mod n) [1..truncate . sqrt $ fromIntegral n]

sumf n = ((subtract n) . sum . factors) n
amicable' x n = (n == sumf x) && (x /= n)
amicable n = amicable' (sumf n) n

main = putStrLn $ show $ sum [n | n <- [2..10000], amicable n]
