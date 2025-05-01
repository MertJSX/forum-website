import { Link } from "react-router";
import { FaHome } from "react-icons/fa";
import { IoTelescope } from "react-icons/io5";
import { FaUser } from "react-icons/fa";
import { FaArrowRightFromBracket } from "react-icons/fa6";

const UserNavbarItems = () => {
  return (
    <div className="bg-gray-800 flex flex-row justify-center items-center pb-1 navbar">
      <h1 className="text-2xl text-blue-300 font-bold italic mr-10 select-none">
        Forum
      </h1>
      <Link
        className="flex justify-center items-center gap-2 text-lg text-white px-4 pt-2 pb-1 navbar-item"
        to="/"
      >
        <FaHome /> Home
      </Link>
      <Link
        className="flex justify-center items-center gap-2 text-lg text-white px-4 pt-2 pb-1 navbar-item"
        to="/explore"
      >
        <IoTelescope /> Explore
      </Link>
      <Link
        className="flex justify-center items-center gap-2 text-lg text-white px-4 pt-2 pb-1 navbar-item"
        to="/profile"
      >
        <FaUser /> Profile
      </Link>
      <Link
        className="flex justify-center items-center gap-2 text-lg text-white px-4 pt-2 pb-1 navbar-item"
        to="/logout"
      >
        <FaArrowRightFromBracket /> Logout
      </Link>
    </div>
  );
};

export default UserNavbarItems;
