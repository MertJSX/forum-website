import "./App.css";
import { Routes, Route, BrowserRouter } from "react-router";
import Home from "./pages/Home/Home";
import SignIn from "./pages/SignIn/SignIn";
import SignUp from "./pages/SignUp/SignUp";
import Logout from "./pages/Logout";
import NotFoundPage from "./pages/NotFoundPage";
import Explore from "./pages/Explore/Explore";
import Profile from "./pages/Profile/Profile";

function App() {
  return (
    <BrowserRouter>
      <Routes>
        <Route index element={<Home />} />
        <Route path="/signin" element={<SignIn />} />
        <Route path="/signup" element={<SignUp />} />
        <Route path="/signup" element={<Explore />} />
        <Route path="/signup" element={<Profile />} />
        <Route path="/logout" element={<Logout />} />
        <Route path="/*" element={<NotFoundPage />} />
      </Routes>
    </BrowserRouter>
  );
}

export default App;
