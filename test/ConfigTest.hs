module ConfigTest (configTests) where

import Test.HUnit ( assertEqual, Test(..) )
import Config
  ( Config (..),
    AllowedApplication (..),
    readConfigFile
  )

buildAllowedApplicationList :: [AllowedApplication]
buildAllowedApplicationList =
  [
    AllowedApplication {
      name = "bhaskara",
      key = "bhaskara-key",
      host = "bhaskara-host"
    },
    AllowedApplication {
      name = "delta",
      key = "delta-key",
      host = "delta-host"
    }
  ]

buildExpectConfig :: Config
buildExpectConfig = Config {
  allowedApplications = buildAllowedApplicationList
}

readConfigFileTest :: Test
readConfigFileTest = TestCase $ do
  result <- readConfigFile "test-data/config-two.json"
  assertEqual "should read a config file" (Right buildExpectConfig) result

configTests :: Test
configTests = TestList [TestLabel "readConfigFileTest" readConfigFileTest]
