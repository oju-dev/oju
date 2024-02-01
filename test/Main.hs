module Main (main) where

import Test.HUnit
  ( runTestTT,
    Counts(errors, failures),
    Test(TestList)
  )
import ConfigTest (configTests)
--import RequesterTest (requesterTests)
import Control.Monad (when)
import qualified System.Exit as Exit

tests :: Test
tests = TestList [configTests]

main :: IO ()
main = do
  result <- runTestTT tests
  when (failures result > 0) $ Exit.exitFailure
  when (errors result > 0) $ Exit.exitFailure
