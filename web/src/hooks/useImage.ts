import useSWR, { mutate } from "swr";
import { produce } from "immer";

import {
  GetImagesDTO,
  ImageAPI,
  ImageResponse,
  SaveImageDTO,
  SearchImagesDTO,
  TImage,
} from "../lib/api";

const api = new ImageAPI();

export type UseImage = {
  response: ImageResponse<TImage[]> | undefined;
  isLoading: boolean;
  isError: any;
};

export const useImages = (key: GetImagesDTO) => {
  const { data, error } = useSWR(["/images", key], (...args) =>
    api.getImages(args[1])
  );

  return {
    response: data,
    isLoading: !error && !data,
    isError: error,
  };
};

export const useSearch = (key: SearchImagesDTO) => {
  const { data, error } = useSWR(["/search", key], (...args) =>
    api.search(args[1])
  );

  return {
    response: data,
    isLoading: !error && !data,
    isError: error,
  };
};

export const useSaveImage = () => {
  const handleSave = async (param: SaveImageDTO) => {
    try {
      const result = await mutate(
        ["/search", { search: "", limit: 9 }],
        api.saveImage(param),
        {
          populateCache: (
            updated: ImageResponse<TImage>,
            current: ImageResponse<TImage[]>
          ) => {
            if(updated.status !== "success") return current

            const newData = produce(current.data, (draft) => {
              draft.unshift(updated.data);
            });

            return {
              ...current,
              data: newData,
            };
          },
          revalidate: false,
        }
      );

      return result;
    } catch (error) {
      console.log({ error });
    }
  };

  return {
    handleSave,
  };
};

export const useDelete = () => {
  const handleDelete = async (id: string) => {
    try {
      const data = await mutate(
        ["/search", { search: "", limit: 9 }],
        api.deleteImage(id),
        {
          populateCache: (
            res: ImageResponse<boolean>,
            current: ImageResponse<TImage[]>
          ) => {
            if (!res.data) current; // failed to delete

            const newData = produce(current.data, (draft) => {
              const index = draft.findIndex((d) => d.id === id);
              if (index !== -1) draft.splice(index, 1);
            });

            return {
              ...current,
              data: newData,
            };
          },
        }
      );

      return data;
    } catch (error) {
      console.log({ error });
    }
  };

  return { handleDelete };
};

export const useHash = () => {
  const hash = async (url:string) =>{
    return await api.hash(url)
  }

  return { hash }
}