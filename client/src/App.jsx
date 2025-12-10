import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import { useState, useEffect } from 'react';
import Login from './pages/Login';
import Register from './pages/Register';
import Dashboard from './pages/Dashboard';
import MasterBarang from './pages/MasterBarang';
import Pembelian from './pages/Pembelian';
import Penjualan from './pages/Penjualan';
import HistoryStok from './pages/HistoryStok';
import Laporan from './pages/Laporan';
import Layout from './components/Layout';

function App() {
  const [isAuthenticated, setIsAuthenticated] = useState(false);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    const token = localStorage.getItem('token');
    setIsAuthenticated(!!token);
    setLoading(false);
  }, []);

  const PrivateRoute = ({ children }) => {
    if (loading) {
      return <div className="min-h-screen flex items-center justify-center">Loading...</div>;
    }
    return isAuthenticated ? children : <Navigate to="/login" />;
  };

  return (
    <Router>
      <Routes>
        <Route path="/login" element={<Login setIsAuthenticated={setIsAuthenticated} />} />
        <Route path="/register" element={<Register />} />
        <Route
          path="/"
          element={
            <PrivateRoute>
              <Layout setIsAuthenticated={setIsAuthenticated}>
                <Dashboard />
              </Layout>
            </PrivateRoute>
          }
        />
        <Route
          path="/barang"
          element={
            <PrivateRoute>
              <Layout setIsAuthenticated={setIsAuthenticated}>
                <MasterBarang />
              </Layout>
            </PrivateRoute>
          }
        />
        <Route
          path="/pembelian"
          element={
            <PrivateRoute>
              <Layout setIsAuthenticated={setIsAuthenticated}>
                <Pembelian />
              </Layout>
            </PrivateRoute>
          }
        />
        <Route
          path="/penjualan"
          element={
            <PrivateRoute>
              <Layout setIsAuthenticated={setIsAuthenticated}>
                <Penjualan />
              </Layout>
            </PrivateRoute>
          }
        />
        <Route
          path="/history-stok"
          element={
            <PrivateRoute>
              <Layout setIsAuthenticated={setIsAuthenticated}>
                <HistoryStok />
              </Layout>
            </PrivateRoute>
          }
        />
        <Route
          path="/laporan"
          element={
            <PrivateRoute>
              <Layout setIsAuthenticated={setIsAuthenticated}>
                <Laporan />
              </Layout>
            </PrivateRoute>
          }
        />
      </Routes>
    </Router>
  );
}

export default App;
