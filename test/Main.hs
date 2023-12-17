module Main (main) where

import Test.HUnit
import ConfigTest (configTests)
import qualified System.Exit as Exit
import Control.Monad (when)

main :: IO ()
main = do
  result <- runTestTT configTests
  when (failures result > 0) $ Exit.exitFailure
  when (errors result > 0) $ Exit.exitFailure
  Exit.exitSuccess
