import GuestNavbarItems from "./GuestNavbarItems";
import UserNavbarItems from "./UserNavbarItems";
import Cookies from "js-cookie";
import "./Navbar.css";

const Navbar = () => {
  return (
    <div>
      {
        Cookies.get("token") ?
        <UserNavbarItems /> :
        <GuestNavbarItems />
      }
    </div>
  );
};

export default Navbar;
