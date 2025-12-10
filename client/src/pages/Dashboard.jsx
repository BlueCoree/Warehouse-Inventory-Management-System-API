import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { stokService, laporanService } from '../api/services';

function Dashboard() {
  const [stats, setStats] = useState({
    totalBarang: 0,
    totalNilaiStok: 0,
    lowStock: 0,
  });
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const [stokRes, laporanRes] = await Promise.all([
        stokService.getAll(),
        laporanService.stok(),
      ]);

      const stokData = stokRes.data.data || [];
      const laporanData = laporanRes.data.data || {};
      
      // Get active items from laporan (excluding deleted)
      const laporanItems = laporanData.data || [];
      const activeItems = laporanItems.filter((item) => !item.deleted);
      const lowStockCount = activeItems.filter((item) => item.stok_akhir < 10).length;

      setStats({
        totalBarang: laporanData.total_items || 0,
        totalNilaiStok: laporanData.total_nilai || 0,
        lowStock: lowStockCount,
      });
    } catch (error) {
      console.error('Error fetching dashboard data:', error);
    } finally {
      setLoading(false);
    }
  };

  if (loading) {
    return <div className="text-center py-10">Loading...</div>;
  }

  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">Dashboard</h1>

      <div className="grid grid-cols-1 md:grid-cols-3 gap-6 mb-8">
        <div className="bg-white p-6 rounded-lg shadow">
          <h3 className="text-gray-500 text-sm font-medium">Total Barang</h3>
          <p className="text-3xl font-bold mt-2">{stats.totalBarang}</p>
        </div>

        <div className="bg-white p-6 rounded-lg shadow">
          <h3 className="text-gray-500 text-sm font-medium">Total Nilai Stok</h3>
          <p className="text-3xl font-bold mt-2">
            Rp {stats.totalNilaiStok.toLocaleString('id-ID')}
          </p>
        </div>

        <div className="bg-white p-6 rounded-lg shadow">
          <h3 className="text-gray-500 text-sm font-medium">Stok Rendah (&lt;10)</h3>
          <p className="text-3xl font-bold mt-2 text-red-600">{stats.lowStock}</p>
        </div>
      </div>

      <div className="bg-white p-6 rounded-lg shadow">
        <h2 className="text-xl font-bold mb-4">Quick Access</h2>
        <div className="grid grid-cols-2 md:grid-cols-4 gap-4">
          <Link
            to="/barang"
            className="bg-blue-600 hover:bg-blue-700 text-white p-4 rounded-lg text-center"
          >
            Master Barang
          </Link>
          <Link
            to="/pembelian"
            className="bg-green-600 hover:bg-green-700 text-white p-4 rounded-lg text-center"
          >
            Pembelian
          </Link>
          <Link
            to="/penjualan"
            className="bg-purple-600 hover:bg-purple-700 text-white p-4 rounded-lg text-center"
          >
            Penjualan
          </Link>
          <Link
            to="/laporan"
            className="bg-orange-600 hover:bg-orange-700 text-white p-4 rounded-lg text-center"
          >
            Laporan
          </Link>
        </div>
      </div>
    </div>
  );
}

export default Dashboard;
