import mongoose from "../db.config";

const userSchema = new mongoose.Schema({
  name: String,
  username: String,
  email: String,
});

const User = mongoose.model("User", userSchema);

const createUser = async () => {
  const newUser = new User({
    name: "Vivial",
    username: "vivial",
    email: "vivial@me.com",
  });

  try {
    const savedUser = await newUser.save();
    console.log('User saved:', savedUser);
  } catch (err) {
    console.error('Error saving movie:', err);
  }
};

createUser();

const updateMovie = async (username: String, email: String) => {
  try {
    const updatedMovie = await User.findByIdAndUpdate(username, email).exec();
    console.log('Updated movie:', updatedMovie);
  } catch (err) {
    console.error('Error updating movie:', err);
  }
};

updateMovie('some-movie-id');

const getMovies = async () => {
  try {
    const movies = await User.find({}).exec();
    console.log('Movies:', movies);
  } catch (err) {
    console.error('Error fetching movies:', err);
  }
};

getMovies();
