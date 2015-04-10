-- Problem 4: Largest palindrome product
--
-- A palindromic number reads the same both ways. 
-- The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
--
-- Find the largest palindrome made from the product of two 3-digit numbers.

palindrome x = (reverse $ show x) == (show x)

main = putStrLn $ show $ maximum $ filter palindrome [x * y | x <- [100..999], y <- [100..999]]
