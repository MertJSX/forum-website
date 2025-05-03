import { useEffect, useState } from "react";
import Navbar from "../../components/Navbar/Navbar"
import Posts from "../../components/Posts/Posts"
import axios from "axios";

const Explore = () => {
  const [posts, setPosts] = useState<any>([]);
  function getPosts() {
    axios.get("/api/posts")
      .then((res) => {
        console.log(res.data);
        setPosts(res.data);
      })
      .catch((err) => {
        console.log(err);
      }
      );
  }
  useEffect(() => {
    getPosts();
  }, []);
  return (
    <div>
        <Navbar />
        <h1 className="text-white text-2xl text-center mt-10">Explore newest posts!</h1>
        <Posts posts={posts} />
    </div>
  )
}

export default Explore