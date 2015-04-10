-- Problem 3
--
-- The prime factors of 13195 are 5, 7, 13 and 29.
--
-- What is the largest prime factor of the number 600851475143 ?

findFactor n = findFactor' n 2

findFactor' n f
    | n `mod` f == 0 = findFactor' (n `div` f) 2
    | (n `mod` f /= 0) && (f * 2 < n) = findFactor' n (f + 1)
    | (n `mod` f /= 0) && (f * 2 >= n) = n

main = putStrLn $ show $ findFactor 600851475143
