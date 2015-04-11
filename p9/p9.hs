-- Problem 9: Special Pythagorean triplet
--
-- https://projecteuler.net/problem=9
--
-- A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,
-- a2 + b2 = c2
--
-- For example, 3^2 + 4^2 = 9 + 16 = 25 = 5^2.
--
-- There exists exactly one Pythagorean triplet for which a + b + c = 1000.
-- Find the product abc. 

-- FIXME - 2slow...

main = putStrLn $ show $ take 1 [a * b * c | a <- reverse [1..1000], b <- reverse [1..1000], c <- reverse [1..1000], c > b, b > a, a + b + c == 1000, a^2 + b^2 == c^2]
