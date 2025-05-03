import "./App.css";
import { Routes, Route, BrowserRouter } from "react-router";
import Home from "./pages/Home/Home";
import SignIn from "./pages/SignIn/SignIn";
import SignUp from "./pages/SignUp/SignUp";
import Logout from "./pages/Logout";
import NotFoundPage from "./pages/NotFoundPage";
import Explore from "./pages/Explore/Explore";
import Profile from "./pages/Profile/Profile";
import Post from "./pages/Post/Post";
import EditPost from "./pages/EditPost/EditPost";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route index element={<Home />} />
        <Route path="/signin" element={<SignIn />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/explore" element={<Explore />} />
        <Route path="/profile" element={<Profile />} />
        <Route path="/profile/:userid" element={<Profile />} />
        <Route path="/post/:postid" element={<Post />} />
        <Route path="/editpost/:postid" element={<EditPost />} />
        <Route path="/logout" element={<Logout />} />
        <Route path="/*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
