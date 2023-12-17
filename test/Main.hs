module Main (main) where

import Test.HUnit
import ConfigTest (configTests)
import RequesterTest (requesterTests)
import qualified System.Exit as Exit
import Control.Monad (when)

main :: IO ()
main = do
  result <- runTestTT $ requesterTests
  when (failures result > 0) $ Exit.exitFailure
  when (errors result > 0) $ Exit.exitFailure
  Exit.exitSuccess
