import { Box } from "@chakra-ui/react";
import { createContext, useState } from "react";
import Images from "./components/Images";
import Navigation from "./components/Navigation";
import { dummy } from "./dummy";
import { TImage } from "./lib/api";

export interface IAppContext {
  data: TImage[];
  handleChangeData: (q: string) => void;
}

export const AppContext = createContext<IAppContext>({
  data: [],
  handleChangeData: () => {},
});

export const AppContextProvider = AppContext.Provider;

function App() {
  const [data, setData] = useState(dummy);

  const handleChangeData = (q: string) => {
    const filter = dummy.filter(({ label }) =>
      label.toLowerCase().includes(q.toLowerCase())
    );

    setData(filter);
  };

  return (
    <Box>
      <AppContextProvider value={{ data, handleChangeData }}>
        <Navigation />
        <Images />
      </AppContextProvider>
    </Box>
  );
}

export default App;
