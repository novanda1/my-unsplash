import useSWR, { KeyedMutator, mutate } from "swr";
import useSWRInfinite from "swr/infinite";
import { produce } from "immer";

import {
  GetImagesDTO,
  ImageAPI,
  ImageResponse,
  SaveImageDTO,
  SearchImagesDTO,
  TImage,
} from "../lib/api";
import { useContext } from "react";
import { ImagesContext } from "../context/app";

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
  const res = useSWRInfinite(
    (index, prevPage: ImageResponse<TImage[]>) => {
      // reached the end
      if (prevPage && !prevPage.data?.length) return null;

      // first page, we don't have `prevPage`
      if (index === 0) return ["/search", key];

      // add the cursor to the param
      const nextCursor =
        prevPage.data && prevPage.data[prevPage.data.length - 1].id;

      return ["/search", { ...key, cursor: nextCursor }];
    },
    (...args) => api.search(args[1])
  );

  return {
    ...res,
    response: res.data,
    isLoading: !res.error && !res.data,
    isError: res.error,
  };
};

export const useImageActions = () => {
  const { mutate: _mutate } = useContext(ImagesContext);

  const handleSave = async (param: SaveImageDTO) => {
    let saveResponse = await api.saveImage(param);

    if (saveResponse.status === "error") return;
    (_mutate.mutate as KeyedMutator<ImageResponse<TImage[]>[]>)(
      (old) => {
        if (!old?.length) return;
        const newData = old;
        newData[0].data.unshift(saveResponse.data);

        return newData;
      },
      { revalidate: false }
    );
  };

  const handleDelete = async (id: string) => {
    if (typeof _mutate.mutate !== "function") return;

    const deleteReq = await api.deleteImage(id);
    if (!deleteReq.data) return;

    try {
      let data;
      await (_mutate.mutate as KeyedMutator<ImageResponse<TImage[]>[]>)(
        (data) => {
          const deleted = data?.map((d) => ({
            ...d,
            data: d.data.filter((img) => img.id !== id),
          }));

          data = deleted;

          return deleted;
        },
        { revalidate: false }
      );

      return data;
    } catch (error) {
      console.log({ error });
    }
  };

  const hash = async (url: string) => {
    return await api.hash(url);
  };

  return {
    handleSave,
    handleDelete,
    hash,
  };
};
