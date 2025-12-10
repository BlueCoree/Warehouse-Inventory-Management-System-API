import { useState, useEffect } from 'react';
import { stokService } from '../api/services';

function HistoryStok() {
  const [history, setHistory] = useState([]);
  const [loading, setLoading] = useState(true);

  useEffect(() => {
    fetchHistory();
  }, []);

  const fetchHistory = async () => {
    try {
      const response = await stokService.getHistory();
      const data = response.data.data?.data || [];
      setHistory(data);
    } catch (error) {
      console.error('Error fetching history:', error);
    } finally {
      setLoading(false);
    }
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleString('id-ID', {
      year: 'numeric',
      month: '2-digit',
      day: '2-digit',
      hour: '2-digit',
      minute: '2-digit',
      second: '2-digit'
    });
  };

  if (loading) {
    return <div className="text-center py-10">Loading...</div>;
  }

  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">History Pergerakan Stok</h1>

      <div className="bg-white rounded-lg shadow overflow-hidden">
        <table className="min-w-full">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Tanggal</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Barang</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Jenis</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Jumlah</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Stok Sebelum</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Stok Sesudah</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Keterangan</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">User</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200">
            {history.map((item) => (
              <tr key={item.id}>
                <td className="px-6 py-4 text-sm whitespace-nowrap">{formatDate(item.created_at)}</td>
                <td className="px-6 py-4">
                  <div className="text-sm font-medium">{item.barang.nama_barang}</div>
                  <div className="text-xs text-gray-500">{item.barang.kode_barang}</div>
                </td>
                <td className="px-6 py-4">
                  <span
                    className={`px-2 py-1 text-xs rounded ${
                      item.jenis_transaksi === 'masuk'
                        ? 'bg-green-100 text-green-800'
                        : 'bg-red-100 text-red-800'
                    }`}
                  >
                    {item.jenis_transaksi}
                  </span>
                </td>
                <td className="px-6 py-4 font-medium">{item.jumlah}</td>
                <td className="px-6 py-4">{item.stok_sebelum}</td>
                <td className="px-6 py-4">{item.stok_sesudah}</td>
                <td className="px-6 py-4 text-sm">{item.keterangan}</td>
                <td className="px-6 py-4 text-sm text-gray-500">{item.user.username}</td>
              </tr>
            ))}
          </tbody>
        </table>
        {history.length === 0 && (
          <div className="text-center py-10 text-gray-500">No history found</div>
        )}
      </div>
    </div>
  );
}

export default HistoryStok;
