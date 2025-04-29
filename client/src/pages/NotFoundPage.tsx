import React from 'react';
import { Link } from 'react-router';

const NotFoundPage: React.FC = () => {
    return (
        <div className="flex flex-col items-center justify-center min-h-screen bg-gray-900 text-white">
            <h1 className="text-6xl font-bold text-blue-300">404</h1>
            <p className="text-xl text-blue-100 mt-4">Page Not Found</p>
            <Link
                to="/"
                className="mt-6 text-lg text-blue-400 hover:text-blue-200 hover:underline hover:text-xl transition"
            >
                Go Back to Home
            </Link>
        </div>
    );
};

export default NotFoundPage;