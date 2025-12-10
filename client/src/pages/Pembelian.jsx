import { useState, useEffect } from 'react';
import { pembelianService, barangService } from '../api/services';
import Swal from 'sweetalert2';
import CountdownTimer from '../components/CountdownTimer';

function Pembelian() {
  const [pembelian, setPembelian] = useState([]);
  const [barang, setBarang] = useState([]);
  const [showModal, setShowModal] = useState(false);
  const [loading, setLoading] = useState(true);
  const [formData, setFormData] = useState({
    no_faktur: '',
    supplier: '',
    details: [],
  });
  const [newDetail, setNewDetail] = useState({
    barang_id: '',
    qty: '',
    harga: '',
  });

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const [pembelianRes, barangRes] = await Promise.all([
        pembelianService.getAll(),
        barangService.getWithStok(),
      ]);
      setPembelian(pembelianRes.data.data || []);
      setBarang(barangRes.data.data || []);
    } catch (error) {
      console.error('Error fetching data:', error);
    } finally {
      setLoading(false);
    }
  };

  const addDetail = () => {
    if (!newDetail.barang_id || !newDetail.qty) {
      Swal.fire({
        icon: 'warning',
        title: 'Oops...',
        text: 'Please select barang and enter quantity',
      });
      return;
    }

    const selectedBarang = barang.find((b) => b.ID === parseInt(newDetail.barang_id));
    const hargaBeli = newDetail.harga || selectedBarang?.harga_beli || 0;
    
    setFormData({
      ...formData,
      details: [
        ...formData.details,
        {
          barang_id: parseInt(newDetail.barang_id),
          qty: parseInt(newDetail.qty),
          harga: parseFloat(hargaBeli),
          barang_nama: selectedBarang?.nama_barang,
        },
      ],
    });
    setNewDetail({ barang_id: '', qty: '', harga: '' });
  };

  const removeDetail = (index) => {
    setFormData({
      ...formData,
      details: formData.details.filter((_, i) => i !== index),
    });
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    if (formData.details.length === 0) {
      Swal.fire({
        icon: 'warning',
        title: 'Oops...',
        text: 'Please add at least one detail',
      });
      return;
    }

    try {
      await pembelianService.create(formData);
      fetchData();
      handleCloseModal();
      Swal.fire({
        icon: 'success',
        title: 'Success!',
        text: 'Pembelian created successfully',
        timer: 2000,
      });
    } catch (error) {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: error.response?.data?.message || 'Error creating pembelian',
      });
    }
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setFormData({ no_faktur: '', supplier: '', details: [] });
    setNewDetail({ barang_id: '', qty: '', harga: '' });
  };

  const handleSelesaikan = async (id) => {
    const result = await Swal.fire({
      title: 'Selesaikan Pembelian?',
      text: 'Stok akan otomatis bertambah setelah diselesaikan.',
      icon: 'question',
      showCancelButton: true,
      confirmButtonColor: '#3085d6',
      cancelButtonColor: '#d33',
      confirmButtonText: 'Ya, Selesaikan!',
      cancelButtonText: 'Batal',
    });

    if (result.isConfirmed) {
      try {
        await pembelianService.selesaikan(id);
        fetchData();
        Swal.fire({
          icon: 'success',
          title: 'Berhasil!',
          text: 'Pembelian berhasil diselesaikan!',
          timer: 2000,
        });
      } catch (error) {
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: error.response?.data?.message || 'Error completing pembelian',
        });
      }
    }
  };

  const formatDate = (dateString) => {
    return new Date(dateString).toLocaleString('id-ID');
  };

  const getCurrentUser = () => {
    const userStr = localStorage.getItem('user');
    return userStr ? JSON.parse(userStr) : null;
  };

  const canSelesaikan = (item) => {
    const currentUser = getCurrentUser();
    return currentUser && item.user_id === currentUser.ID;
  };

  const handleExpired = async (id) => {
    try {
      await pembelianService.cancel(id);
      fetchData();
      Swal.fire({
        icon: 'warning',
        title: 'Transaksi Expired',
        text: 'Pembelian telah dibatalkan karena melewati batas waktu 5 menit.',
        timer: 3000,
      });
    } catch (error) {
      console.error('Error cancelling expired pembelian:', error);
    }
  };

  const getItemsSummary = (details) => {
    if (!details || details.length === 0) return '-';
    return details.map(d => d.Barang?.nama_barang || 'Unknown').join(', ');
  };

  if (loading) {
    return <div className="text-center py-10">Loading...</div>;
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Transaksi Pembelian</h1>
        <button
          onClick={() => setShowModal(true)}
          className="bg-green-600 hover:bg-green-700 text-white px-4 py-2 rounded-md"
        >
          + Buat Pembelian
        </button>
      </div>

      <div className="bg-white rounded-lg shadow overflow-hidden">
        <table className="min-w-full">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">No Faktur</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Supplier</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Items</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Total</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Status</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Waktu Tersisa</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200">
            {pembelian.map((item) => (
              <tr key={item.ID}>
                <td className="px-6 py-4 font-medium">{item.no_faktur}</td>
                <td className="px-6 py-4">{item.supplier}</td>
                <td className="px-6 py-4 text-sm max-w-xs truncate" title={getItemsSummary(item.details)}>
                  {getItemsSummary(item.details)}
                </td>
                <td className="px-6 py-4">Rp {item.total?.toLocaleString('id-ID')}</td>
                <td className="px-6 py-4">
                  <span className={`px-2 py-1 rounded text-sm ${
                    item.status === 'selesai' 
                      ? 'bg-green-100 text-green-800' 
                      : item.status === 'cancel'
                      ? 'bg-red-100 text-red-800'
                      : 'bg-yellow-100 text-yellow-800'
                  }`}>
                    {item.status}
                  </span>
                </td>
                <td className="px-6 py-4">
                  <CountdownTimer 
                    createdAt={item.CreatedAt} 
                    status={item.status}
                    onExpired={() => handleExpired(item.ID)}
                  />
                </td>
                <td className="px-6 py-4">
                  {item.status === 'pending' && canSelesaikan(item) && (
                    <button
                      onClick={() => handleSelesaikan(item.ID)}
                      className="bg-blue-600 hover:bg-blue-700 text-white px-3 py-1 rounded text-sm"
                    >
                      Selesaikan
                    </button>
                  )}
                  {item.status === 'pending' && !canSelesaikan(item) && (
                    <button
                      disabled
                      className="bg-gray-300 text-gray-500 px-3 py-1 rounded text-sm cursor-not-allowed"
                      title="Hanya user yang membuat pembelian yang dapat menyelesaikan"
                    >
                      Selesaikan
                    </button>
                  )}
                  {item.status === 'selesai' && (
                    <span className="text-gray-400 text-sm">Completed</span>
                  )}
                  {item.status === 'cancel' && (
                    <span className="text-red-400 text-sm">Cancelled</span>
                  )}
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {showModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50 overflow-y-auto">
          <div className="bg-white rounded-lg p-6 w-full max-w-4xl m-4">
            <h2 className="text-2xl font-bold mb-4">Buat Pembelian Baru</h2>
            <form onSubmit={handleSubmit}>
              <div className="grid grid-cols-2 gap-4 mb-4">
                <div>
                  <label className="block text-sm font-medium mb-1">No Faktur</label>
                  <input
                    type="text"
                    value={formData.no_faktur}
                    onChange={(e) => setFormData({ ...formData, no_faktur: e.target.value })}
                    className="w-full px-3 py-2 border rounded-md"
                    required
                  />
                </div>
                <div>
                  <label className="block text-sm font-medium mb-1">Supplier</label>
                  <input
                    type="text"
                    value={formData.supplier}
                    onChange={(e) => setFormData({ ...formData, supplier: e.target.value })}
                    className="w-full px-3 py-2 border rounded-md"
                    required
                  />
                </div>
              </div>

              <div className="border-t pt-4 mt-4">
                <h3 className="font-bold mb-3">Detail Items</h3>
                <div className="grid grid-cols-4 gap-2 mb-3">
                  <select
                    value={newDetail.barang_id}
                    onChange={(e) => {
                      const selectedBarang = barang.find((b) => b.ID === parseInt(e.target.value));
                      setNewDetail({ 
                        ...newDetail, 
                        barang_id: e.target.value,
                        harga: selectedBarang?.harga_beli || ''
                      });
                    }}
                    className="px-3 py-2 border rounded-md"
                  >
                    <option value="">Select Barang</option>
                    {barang.map((b) => (
                      <option key={b.ID} value={b.ID}>
                        {b.kode_barang} - {b.nama_barang}
                      </option>
                    ))}
                  </select>
                  <input
                    type="number"
                    placeholder="Qty"
                    value={newDetail.qty}
                    onChange={(e) => setNewDetail({ ...newDetail, qty: e.target.value })}
                    className="px-3 py-2 border rounded-md"
                  />
                  <input
                    type="number"
                    placeholder="Harga"
                    value={newDetail.harga}
                    onChange={(e) => setNewDetail({ ...newDetail, harga: e.target.value })}
                    className="px-3 py-2 border rounded-md"
                  />
                  <button
                    type="button"
                    onClick={addDetail}
                    className="bg-blue-600 text-white px-4 py-2 rounded-md hover:bg-blue-700"
                  >
                    Add
                  </button>
                </div>

                {formData.details.length > 0 && (
                  <table className="w-full text-sm">
                    <thead className="bg-gray-50">
                      <tr>
                        <th className="px-4 py-2 text-left">Barang</th>
                        <th className="px-4 py-2 text-left">Qty</th>
                        <th className="px-4 py-2 text-left">Harga</th>
                        <th className="px-4 py-2 text-left">Subtotal</th>
                        <th className="px-4 py-2"></th>
                      </tr>
                    </thead>
                    <tbody>
                      {formData.details.map((detail, index) => (
                        <tr key={index} className="border-t">
                          <td className="px-4 py-2">{detail.barang_nama}</td>
                          <td className="px-4 py-2">{detail.qty}</td>
                          <td className="px-4 py-2">Rp {detail.harga.toLocaleString('id-ID')}</td>
                          <td className="px-4 py-2">
                            Rp {(detail.qty * detail.harga).toLocaleString('id-ID')}
                          </td>
                          <td className="px-4 py-2">
                            <button
                              type="button"
                              onClick={() => removeDetail(index)}
                              className="text-red-600 hover:text-red-800"
                            >
                              Remove
                            </button>
                          </td>
                        </tr>
                      ))}
                      <tr className="border-t font-bold">
                        <td colSpan="3" className="px-4 py-2 text-right">Total:</td>
                        <td className="px-4 py-2">
                          Rp {formData.details.reduce((sum, d) => sum + (d.qty * d.harga), 0).toLocaleString('id-ID')}
                        </td>
                        <td></td>
                      </tr>
                    </tbody>
                  </table>
                )}
              </div>

              <div className="flex justify-end gap-2 mt-6">
                <button
                  type="button"
                  onClick={handleCloseModal}
                  className="px-4 py-2 border rounded-md hover:bg-gray-50"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 bg-green-600 text-white rounded-md hover:bg-green-700"
                >
                  Simpan
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}

export default Pembelian;
