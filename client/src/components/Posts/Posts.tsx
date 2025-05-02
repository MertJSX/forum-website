import { Link } from "react-router";
import TextTruncate from "../../utils/TextTruncate";

type PostsProps = {
    posts?: any[];
};

const Posts = ({ posts }: PostsProps) => {
  return (
    <div>
        { posts ? (
            <div className="flex flex-col items-center mt-10 w-1/2 justify-center mx-auto">
                {posts.map((post, index) => (
                    <Link target="_blank" to={`/post/${post.id}`} key={index} className="bg-gray-700 hover:bg-gray-600 transition-all hover:cursor-pointer text-white p-4 rounded-lg mb-4 w-full">
                        <div className="flex items-center gap-2">
                            <h2 className="text-xl font-bold">{post.title}</h2>
                            -
                            <h2 className="text-gray-400 hover:text-blue-300 text-lg font-bold italic">{post.author}</h2>
                        </div>
                        <p><TextTruncate text={post.content} /></p>
                    </Link>
                ))}
            </div>
        ) : (
            <h1 className="text-white text-2xl text-center">No posts available</h1>
        )}
    </div>
  )
}

export default Posts