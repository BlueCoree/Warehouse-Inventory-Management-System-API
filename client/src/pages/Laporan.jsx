import { useState, useEffect } from 'react';
import { laporanService } from '../api/services';

function Laporan() {
  const [activeTab, setActiveTab] = useState('stok');
  const [laporanStok, setLaporanStok] = useState([]);
  const [laporanPenjualan, setLaporanPenjualan] = useState([]);
  const [laporanPembelian, setLaporanPembelian] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchLaporan();
  }, [activeTab]);

  const fetchLaporan = async () => {
    setLoading(true);
    try {
      if (activeTab === 'stok') {
        const response = await laporanService.stok();
        setLaporanStok(response.data.data?.data || []);
      } else if (activeTab === 'penjualan') {
        const response = await laporanService.penjualan();
        setLaporanPenjualan(response.data.data?.data || []);
      } else if (activeTab === 'pembelian') {
        const response = await laporanService.pembelian();
        setLaporanPembelian(response.data.data?.data || []);
      }
    } catch (error) {
      console.error('Error fetching laporan:', error);
    } finally {
      setLoading(false);
    }
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleString('id-ID');
  };

  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">Laporan</h1>

      <div className="mb-6 border-b">
        <div className="flex space-x-4">
          <button
            onClick={() => setActiveTab('stok')}
            className={`px-4 py-2 font-medium ${
              activeTab === 'stok'
                ? 'border-b-2 border-blue-600 text-blue-600'
                : 'text-gray-500 hover:text-gray-700'
            }`}
          >
            Laporan Stok
          </button>
          <button
            onClick={() => setActiveTab('penjualan')}
            className={`px-4 py-2 font-medium ${
              activeTab === 'penjualan'
                ? 'border-b-2 border-blue-600 text-blue-600'
                : 'text-gray-500 hover:text-gray-700'
            }`}
          >
            Laporan Penjualan
          </button>
          <button
            onClick={() => setActiveTab('pembelian')}
            className={`px-4 py-2 font-medium ${
              activeTab === 'pembelian'
                ? 'border-b-2 border-blue-600 text-blue-600'
                : 'text-gray-500 hover:text-gray-700'
            }`}
          >
            Laporan Pembelian
          </button>
        </div>
      </div>

      {loading ? (
        <div className="text-center py-10">Loading...</div>
      ) : (
        <>
          {activeTab === 'stok' && (
            <div className="bg-white rounded-lg shadow overflow-hidden">
              <table className="min-w-full">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Kode</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama Barang</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Satuan</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Stok Akhir</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Harga Beli</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nilai Stok</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-200">
                  {laporanStok.map((item) => (
                    <tr key={item.id}>
                      <td className="px-6 py-4">{item.kode_barang}</td>
                      <td className="px-6 py-4">{item.nama_barang}</td>
                      <td className="px-6 py-4">{item.satuan}</td>
                      <td className="px-6 py-4">{item.stok_akhir}</td>
                      <td className="px-6 py-4">Rp {item.harga_beli?.toLocaleString('id-ID')}</td>
                      <td className="px-6 py-4 font-medium">
                        Rp {item.nilai_stok?.toLocaleString('id-ID')}
                      </td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}

          {activeTab === 'penjualan' && (
            <div className="bg-white rounded-lg shadow overflow-hidden">
              <table className="min-w-full">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">No Faktur</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Customer</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Items</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Total</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Tanggal</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">User</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-200">
                  {laporanPenjualan.map((item) => (
                    <tr key={item.id}>
                      <td className="px-6 py-4 font-medium">{item.no_faktur}</td>
                      <td className="px-6 py-4">{item.customer}</td>
                      <td className="px-6 py-4 text-sm text-gray-600">
                        {item.details && item.details.length > 0 ? (
                          <div className="space-y-1">
                            {item.details.map((d, idx) => (
                              <div key={idx}>
                                {d.nama_barang} ({d.qty}x)
                              </div>
                            ))}
                          </div>
                        ) : (
                          <span className="text-gray-400">-</span>
                        )}
                      </td>
                      <td className="px-6 py-4">Rp {item.total?.toLocaleString('id-ID')}</td>
                      <td className="px-6 py-4">
                        <span className="px-2 py-1 bg-purple-100 text-purple-800 rounded text-xs">
                          {item.status}
                        </span>
                      </td>
                      <td className="px-6 py-4 text-sm">{formatDate(item.created_at)}</td>
                      <td className="px-6 py-4 text-sm text-gray-500">{item.username}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}

          {activeTab === 'pembelian' && (
            <div className="bg-white rounded-lg shadow overflow-hidden">
              <table className="min-w-full">
                <thead className="bg-gray-50">
                  <tr>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">No Faktur</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Supplier</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Items</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Total</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Tanggal</th>
                    <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">User</th>
                  </tr>
                </thead>
                <tbody className="divide-y divide-gray-200">
                  {laporanPembelian.map((item) => (
                    <tr key={item.id}>
                      <td className="px-6 py-4 font-medium">{item.no_faktur}</td>
                      <td className="px-6 py-4">{item.supplier}</td>
                      <td className="px-6 py-4 text-sm text-gray-600">
                        {item.details && item.details.length > 0 ? (
                          <div className="space-y-1">
                            {item.details.map((d, idx) => (
                              <div key={idx}>
                                {d.nama_barang} ({d.qty}x)
                              </div>
                            ))}
                          </div>
                        ) : (
                          <span className="text-gray-400">-</span>
                        )}
                      </td>
                      <td className="px-6 py-4">Rp {item.total?.toLocaleString('id-ID')}</td>
                      <td className="px-6 py-4">
                        <span className={`px-2 py-1 rounded text-xs ${
                          item.status === 'selesai' 
                            ? 'bg-green-100 text-green-800' 
                            : 'bg-yellow-100 text-yellow-800'
                        }`}>
                          {item.status}
                        </span>
                      </td>
                      <td className="px-6 py-4 text-sm">{formatDate(item.created_at)}</td>
                      <td className="px-6 py-4 text-sm text-gray-500">{item.username}</td>
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </>
      )}
    </div>
  );
}

export default Laporan;
