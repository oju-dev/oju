module RequesterTest () where

import Data.ByteString.Lazy.Char8 as C

generateByteStringPacket :: ByteString
generateByteStringPacket = C.pack "TRACE 3FAFCF87-BF66-4DC5-84C1-34E178FF55CC AWO\n{\"name\":\"span-name\",\"service\":\"\",\"attributes\":{\"http.url\":\"http://products.api.svc.cluster.local\",\"http.method\":\"POST\",\"http.body.email\":\"test@email.com\"}}"

-- Create a function that will break the packet in two and decode in a Header and a Trace
