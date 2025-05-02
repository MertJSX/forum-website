import { useEffect } from "react";
import Navbar from "../../components/Navbar/Navbar";
import { useState } from "react";
import { FaUserAlt } from "react-icons/fa";
import GetProfile from "../../utils/getProfile";
import { useNavigate } from "react-router";

const Profile = () => {
  const navigate = useNavigate();
  let [userData, setUserData] = useState<any>(null);
  useEffect(() => {
    GetProfile.then((res) => {
        setUserData(res);
    }
    ).catch((err: string) => {
      console.log(err);
      if (err === "Unauthorized") {
        navigate("/signin");
      }
    }
    );
  }, []);
  return (
    <div>
      <Navbar />
      <div className="flex flex-col items-center">
        {userData ? (
          <div className="text-white mt-10">
            <div className="flex justify-center items-center">
              <FaUserAlt className="text-9xl bg-neutral-600 p-5 rounded-4xl"/>
            </div>
            <h1 className="text-center text-white text-3xl">{userData.username}</h1>
            <h2 className="text-center text-white text-xl">{userData.email}</h2>
          </div>
        ) : null}
      </div>
    </div>
  );
};

export default Profile;

