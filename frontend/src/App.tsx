import {BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import './App.css'
import Header from './components/Header';
import LoginPage from './pages/LoginPage';
import RegisterPage from './pages/RegisterPage';
import ProfilePage from './pages/ProfilePage';
import Home from './pages/Home';
import AuthProvider from './components/AuthContext';
import Layout from './components/Layout';
import CarPage from './pages/CarPage';



function App() {
  return (
    <Router>
      <AuthProvider>
        <Routes>
          <Route element={<Layout />}>
            <Route path="/" element={<Home />} />
            <Route path="/cars/:id" element={<CarPage />} />
          </Route>
          <Route path="/register" element={<RegisterPage />} />
          <Route path="/login" element={<LoginPage />} />
          <Route path="/profile" element={<ProfilePage />} />
        </Routes>
      </AuthProvider>
    </Router>
  );
}

export default App;
