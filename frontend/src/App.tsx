import React, { useState } from 'react';
import { Routes, Route, Link } from 'react-router-dom';
import './App.css'
import Header from './components/Header';
import CarList from './components/CarList';
import CarFilter from './components/CarFilter';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import ProfilePage from './pages/ProfilePage';

const HomePage: React.FC = () => {
  const [searchTerm, setSearchTerm] = useState('');
  return (
    <>
      <CarFilter searchTerm={searchTerm} onSearchTermChange={setSearchTerm} />
      <CarList searchTerm={searchTerm} />
    </>
  );
};

function App() {
  const isAuthenticated = false;

  return (
    <div className="container mx-auto p-4">
      <Header />
      <nav className="mb-4">
        <ul className="flex space-x-4 justify-center">
          <li><Link to="/" className="text-blue-600 hover:text-blue-800">Главная</Link></li>
          {isAuthenticated ? (
            <li><Link to="/profile" className="text-blue-600 hover:text-blue-800">Личный кабинет</Link></li>
          ) : (
            <>
              <li><Link to="/login" className="text-blue-600 hover:text-blue-800">Вход</Link></li>
              <li><Link to="/register" className="text-blue-600 hover:text-blue-800">Регистрация</Link></li>
            </>
          )}
        </ul>
      </nav>

      <main>
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/profile" element={<ProfilePage />} />
        </Routes>
      </main>

    </div>
  );
}

export default App;
