module Requester.Request () where

data RequestError =
  ApplicationNotAllowed
    | MalformedHeader
    | MalformedPacket
    | VerbNotAllowed
  deriving (Enum)

data Request = Request
  {
    requestHeader :: Header
  }
