import { Link, useNavigate, useLocation } from 'react-router-dom';

function Layout({ children, setIsAuthenticated }) {
  const navigate = useNavigate();
  const location = useLocation();

  const handleLogout = () => {
    localStorage.removeItem('token');
    localStorage.removeItem('user');
    setIsAuthenticated(false);
    navigate('/login');
  };

  const isActive = (path) => location.pathname === path;

  return (
    <div className="min-h-screen bg-gray-100">
      <nav className="bg-gray-800 text-white">
        <div className="container mx-auto px-4">
          <div className="flex items-center justify-between h-16">
            <div className="flex items-center space-x-4">
              <Link to="/" className="font-bold text-xl">
                Warehouse System
              </Link>
              <div className="hidden md:flex space-x-2">
                <Link
                  to="/"
                  className={`px-3 py-2 rounded-md text-sm font-medium ${
                    isActive('/') ? 'bg-gray-900' : 'hover:bg-gray-700'
                  }`}
                >
                  Dashboard
                </Link>
                <Link
                  to="/barang"
                  className={`px-3 py-2 rounded-md text-sm font-medium ${
                    isActive('/barang') ? 'bg-gray-900' : 'hover:bg-gray-700'
                  }`}
                >
                  Master Barang
                </Link>
                <Link
                  to="/pembelian"
                  className={`px-3 py-2 rounded-md text-sm font-medium ${
                    isActive('/pembelian') ? 'bg-gray-900' : 'hover:bg-gray-700'
                  }`}
                >
                  Pembelian
                </Link>
                <Link
                  to="/penjualan"
                  className={`px-3 py-2 rounded-md text-sm font-medium ${
                    isActive('/penjualan') ? 'bg-gray-900' : 'hover:bg-gray-700'
                  }`}
                >
                  Penjualan
                </Link>
                <Link
                  to="/history-stok"
                  className={`px-3 py-2 rounded-md text-sm font-medium ${
                    isActive('/history-stok') ? 'bg-gray-900' : 'hover:bg-gray-700'
                  }`}
                >
                  History Stok
                </Link>
                <Link
                  to="/laporan"
                  className={`px-3 py-2 rounded-md text-sm font-medium ${
                    isActive('/laporan') ? 'bg-gray-900' : 'hover:bg-gray-700'
                  }`}
                >
                  Laporan
                </Link>
              </div>
            </div>
            <button
              onClick={handleLogout}
              className="bg-red-600 hover:bg-red-700 px-4 py-2 rounded-md text-sm font-medium"
            >
              Logout
            </button>
          </div>
        </div>
      </nav>
      <main className="container mx-auto px-4 py-8">{children}</main>
    </div>
  );
}

export default Layout;
