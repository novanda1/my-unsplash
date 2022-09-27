import { Box, Container } from "@chakra-ui/react";
import { createContext, useState } from "react";
import Images from "./components/Images";
import Navigation from "./components/Navigation";
import { dummy } from "./dummy";
import { useSearch } from "./hooks/useImage";
import { TImage } from "./lib/api";

export interface IAppContext {
  data?: TImage[];
  handleChangeData: (q: string) => void;
}

export const AppContext = createContext<IAppContext>({
  data: [],
  handleChangeData: () => {},
});

export const AppContextProvider = AppContext.Provider;

function App() {
  const [search, setSearch] = useState("");

  const { images, isError, isLoading } = useSearch({ search, limit: 9 });

  const handleChangeData = (q: string) => {
    if (q.length >= 3 || !q.length) setSearch(q.toLowerCase());
  };

  return (
    <Box>
      <AppContextProvider value={{ data: images?.data, handleChangeData }}>
        <Navigation />

        <Container maxW="container.xl" pb={50}>
          {isLoading ? (
            "Loading..."
          ) : isError ? (
            "Something went wrong..."
          ) : (
            <Images />
          )}
        </Container>
      </AppContextProvider>
    </Box>
  );
}

export default App;
