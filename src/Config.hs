module Config
  (
    Config (..),
    AllowedApplication (..),
    readConfigFile
  ) where

import Data.Text (Text)
import Data.Aeson.Casing (snakeCase, aesonDrop)
import Data.Aeson
  (FromJSON (parseJSON),
    eitherDecode,
    genericParseJSON,
  )
import GHC.Generics (Generic)
import Data.ByteString.Lazy as BL ( readFile )

data AllowedApplication = AllowedApplication
        {       name :: !Text,
                key :: !Text,
                host :: Text
        } deriving (Eq, Show, Generic, FromJSON)

newtype Config = Config
        {
                allowedApplications :: [AllowedApplication]
        } deriving (Eq, Show, Generic)

instance FromJSON Config where
  parseJSON = genericParseJSON $ aesonDrop 0 snakeCase

readConfigFile :: FilePath -> IO (Either String Config)
readConfigFile file = do
        contents <- BL.readFile file
        case eitherDecode contents of
                Left err -> pure $ Left err
                Right cfg -> pure $ Right cfg
