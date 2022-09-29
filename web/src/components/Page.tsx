import { Box } from "@chakra-ui/react";
import { PropsWithChildren, useState } from "react";
import { KeyedMutator } from "swr";
import { ImagesContextProvider } from "../context/app";
import { ImageResponse, TImage } from "../lib/api";
import Navigation from "./Navigation";

const Page: React.FC<PropsWithChildren> = ({ children }) => {
  const [query, setQuery] = useState("");
  const [mutate, setMutate] = useState<KeyedMutator<
    ImageResponse<TImage[]>[]
  > | null>(null);

  return (
    <ImagesContextProvider value={{ query, setQuery, mutate, setMutate }}>
      <Box>
        <Navigation />
        {children}
      </Box>
    </ImagesContextProvider>
  );
};

export default Page;
