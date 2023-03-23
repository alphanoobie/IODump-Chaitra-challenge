-- day 1 (haskell)
-- get minimum and maximum of 5 numbers

import Control.Monad (replicateM)
import Control.Monad.State
  ( MonadIO (liftIO),
    MonadState (get),
    StateT,
    evalStateT,
    modify,
  )

getInp :: [String] -> StateT Int IO [String]
getInp xs = do
  x <- get
  if x == 5
    then do
      return xs
    else do
      modify (+ 1)
      x <- liftIO getLine
      getInp (x : xs)

main :: IO ()
main = do
  x <- (evalStateT $ getInp []) 0
  let nums = (read <$> x) :: [Int]
  putStrLn $ "Maximum: " ++ show (maximum nums)
  putStrLn $ "Minumum: " ++ show (minimum nums)
