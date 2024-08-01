import db from "../db.config";

const getMovies = async () => {
  const result = await db
    .collection("movies")
    .findOne({ title: "The Matrix" });
  return result;
};

export { getMovies };
