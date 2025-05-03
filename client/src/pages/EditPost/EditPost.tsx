import React, { useEffect, useState } from 'react';
import { useParams } from 'react-router';
import axios from 'axios';
import Navbar from '../../components/Navbar/Navbar';
import Cookies from 'js-cookie';
import { useNavigate } from 'react-router';

const EditPost = () => {
    const navigate = useNavigate();
    const { postid } = useParams<{ postid: string }>();
    const [post, setPost] = useState<any>(null);

    useEffect(() => {
        const fetchPost = async () => {
            try {
                const response = await axios.get(`/api/posts/${postid}`);
                setPost(response.data);
            } catch (error) {
                console.error('Error fetching post:', error);
            }
        };

        fetchPost();
    }, [postid]);

    const handleInputChange = (e: React.ChangeEvent<HTMLInputElement | HTMLTextAreaElement>) => {
        if (post) {
            setPost({ ...post, [e.target.name]: e.target.value });
        }
    };

    const handleSubmit = async (e: React.FormEvent) => {
        e.preventDefault();
        try {
            await axios.put(`/api/posts/${postid}`, post, {
                headers: { 'Content-Type': 'application/json',
                authorization: Cookies.get('token')},
                
            });
            alert('Post updated successfully!');
            navigate(`/profile/${post.userid ? post.userid : ''}`);
        } catch (error) {
            console.error('Error updating post:', error);
        }
    };

    if (!post) {
        return <div className="text-center text-blue-300">Loading...</div>;
    }

    return (
        <div>
        <Navbar />
        <div className="max-w-2xl mx-auto p-6 bg-gray-800 shadow-md rounded-md mt-20">
            <h1 className="text-3xl text-center font-bold mb-4 text-blue-300">Edit Post</h1>
            <form onSubmit={handleSubmit} className="space-y-4">
                <div>
                    <label htmlFor="title" className="block text-lg font-medium text-blue-200">
                        Title:
                    </label>
                    <input
                        type="text"
                        id="title"
                        name="title"
                        value={post.title}
                        onChange={handleInputChange}
                        className="mt-1 block w-full px-3 py-2 border rounded-md shadow-sm bg-gray-600 text-blue-100 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-lg"
                    />
                </div>
                <div>
                    <label htmlFor="content" className="block text-lg font-medium text-blue-200">
                        Content:
                    </label>
                    <textarea
                        id="content"
                        name="content"
                        value={post.content}
                        onChange={handleInputChange}
                        className="mt-1 block w-full px-3 py-2 border border-gray-600 rounded-md shadow-sm bg-gray-600 text-blue-100 focus:outline-none focus:ring-blue-500 focus:border-blue-500 sm:text-lg"
                        rows={6}
                    />
                </div>
                <button
                    type="submit"
                    className="w-full bg-blue-600 text-white cursor-pointer py-2 px-4 rounded-md hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 text-lg"
                >
                    Save Changes
                </button>
            </form>
        </div>
        </div>
    );
};

export default EditPost;
