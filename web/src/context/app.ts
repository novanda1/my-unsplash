import { createContext } from "react";
import { KeyedMutator } from "swr";
import { ImageResponse, TImage } from "../lib/api";

export interface IImagesContext {
  query: string;
  setQuery: (q: string) => void;
  mutate: any | null;
  setMutate: (q: any | null) => void;
}

export const ImagesContext = createContext<IImagesContext>({
  query: "",
  setQuery: (q: string) => {},
  mutate: null,
  setMutate: () => {},
});

export const ImagesContextProvider = ImagesContext.Provider;
