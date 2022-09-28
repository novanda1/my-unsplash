import { NextPage } from "next";
import { Box } from "@chakra-ui/react";
import { ImagesContextProvider } from "../context/app";
import Navigation from "./Navigation";
import { PropsWithChildren, useState } from "react";
import { useSearch } from "../hooks/useImage";

const Page: React.FC<PropsWithChildren> = ({ children }) => {
  const [search, setSearch] = useState("");
  const data = useSearch({ search, limit: 9 });

  const handleChangeData = (q: string) => {
    if (q.length >= 3 || !q.length) setSearch(q.toLowerCase());
  };

  return (
    <ImagesContextProvider value={{ handleChangeData, images: data }}>
      <Box>
        <Navigation />
        {children}
      </Box>
    </ImagesContextProvider>
  );
};

export default Page;
