module Requester () where

import Data.ByteString.Lazy (ByteString)

handlePacket :: ByteString -> Request
handlePacket packet =
