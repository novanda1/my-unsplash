// Next.js API route support: https://nextjs.org/docs/api-routes/introduction
import type {
  NextApiRequest,
  NextApiResponse,
} from "next";

type Data = {
  name: string;
};

const sleep = () =>
  new Promise((resolve) =>
    setTimeout(resolve, 3000)
  );

export default async function handler(
  req: NextApiRequest,
  res: NextApiResponse<Data>
) {
  await sleep();
  res.status(200).json({ name: "John Doe" });
}
