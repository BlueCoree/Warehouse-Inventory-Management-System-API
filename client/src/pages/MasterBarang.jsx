import { useState, useEffect } from 'react';
import { barangService } from '../api/services';
import Swal from 'sweetalert2';

function MasterBarang() {
  const [barang, setBarang] = useState([]);
  const [loading, setLoading] = useState(true);
  const [showModal, setShowModal] = useState(false);
  const [editMode, setEditMode] = useState(false);
  const [currentId, setCurrentId] = useState(null);
  const [search, setSearch] = useState('');
  const [formData, setFormData] = useState({
    kode_barang: '',
    nama_barang: '',
    deskripsi: '',
    satuan: '',
    harga_beli: '',
    harga_jual: '',
    stok_akhir: '',
  });

  useEffect(() => {
    fetchBarang();
  }, [search]);

  const fetchBarang = async () => {
    try {
      const response = await barangService.getWithStok();
      let data = response.data.data || [];
      
      if (search) {
        data = data.filter(b => 
          b.kode_barang.toLowerCase().includes(search.toLowerCase()) ||
          b.nama_barang.toLowerCase().includes(search.toLowerCase())
        );
      }
      
      setBarang(data);
    } catch (error) {
      console.error('Error fetching barang:', error);
    } finally {
      setLoading(false);
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      const payload = {
        kode_barang: formData.kode_barang,
        nama_barang: formData.nama_barang,
        deskripsi: formData.deskripsi,
        satuan: formData.satuan,
        harga_beli: parseFloat(formData.harga_beli) || 0,
        harga_jual: parseFloat(formData.harga_jual) || 0,
        stok_akhir: parseInt(formData.stok_akhir) || 0,
      };
      
      if (editMode) {
        await barangService.update(currentId, payload);
      } else {
        await barangService.create(payload);
      }
      fetchBarang();
      handleCloseModal();
    } catch (error) {
      Swal.fire({
        icon: 'error',
        title: 'Error',
        text: error.response?.data?.message || 'Error saving data',
      });
    }
  };

  const handleEdit = (item) => {
    setEditMode(true);
    setCurrentId(item.ID);
    setFormData({
      kode_barang: item.kode_barang,
      nama_barang: item.nama_barang,
      deskripsi: item.deskripsi || '',
      satuan: item.satuan || '',
      harga_beli: item.harga_beli || '',
      harga_jual: item.harga_jual || '',
      stok_akhir: item.stok_akhir || 0,
    });
    setShowModal(true);
  };

  const handleDelete = async (id) => {
    const result = await Swal.fire({
      title: 'Are you sure?',
      text: 'You won\'t be able to revert this!',
      icon: 'warning',
      showCancelButton: true,
      confirmButtonColor: '#d33',
      cancelButtonColor: '#3085d6',
      confirmButtonText: 'Yes, delete it!',
      cancelButtonText: 'Cancel',
    });

    if (result.isConfirmed) {
      try {
        await barangService.delete(id);
        fetchBarang();
        Swal.fire({
          icon: 'success',
          title: 'Deleted!',
          text: 'Data has been deleted.',
          timer: 2000,
        });
      } catch (error) {
        Swal.fire({
          icon: 'error',
          title: 'Error',
          text: error.response?.data?.message || 'Error deleting data',
        });
      }
    }
  };

  const handleCloseModal = () => {
    setShowModal(false);
    setEditMode(false);
    setCurrentId(null);
    setFormData({
      kode_barang: '',
      nama_barang: '',
      deskripsi: '',
      satuan: '',
      harga_beli: '',
      harga_jual: '',
      stok_akhir: '',
    });
  };

  if (loading) {
    return <div className="text-center py-10">Loading...</div>;
  }

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <h1 className="text-3xl font-bold">Master Barang</h1>
        <button
          onClick={() => setShowModal(true)}
          className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-md"
        >
          + Tambah Barang
        </button>
      </div>

      <div className="mb-4">
        <input
          type="text"
          placeholder="Search by kode or nama barang..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="w-full md:w-96 px-4 py-2 border border-gray-300 rounded-md"
        />
      </div>

      <div className="bg-white rounded-lg shadow overflow-hidden">
        <table className="min-w-full">
          <thead className="bg-gray-50">
            <tr>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Kode</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Nama Barang</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Satuan</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Harga Beli</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Harga Jual</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Stok</th>
              <th className="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase">Actions</th>
            </tr>
          </thead>
          <tbody className="divide-y divide-gray-200">
            {barang.map((item) => (
              <tr key={item.ID}>
                <td className="px-6 py-4">{item.kode_barang}</td>
                <td className="px-6 py-4">{item.nama_barang}</td>
                <td className="px-6 py-4">{item.satuan}</td>
                <td className="px-6 py-4">Rp {item.harga_beli?.toLocaleString('id-ID')}</td>
                <td className="px-6 py-4">Rp {item.harga_jual?.toLocaleString('id-ID')}</td>
                <td className="px-6 py-4">
                  <span className={`px-2 py-1 rounded ${item.stok_akhir < 10 ? 'bg-red-100 text-red-800' : 'bg-green-100 text-green-800'}`}>
                    {item.stok_akhir}
                  </span>
                </td>
                <td className="px-6 py-4">
                  <button
                    onClick={() => handleEdit(item)}
                    className="text-blue-600 hover:text-blue-800 mr-3"
                  >
                    Edit
                  </button>
                  <button
                    onClick={() => handleDelete(item.ID)}
                    className="text-red-600 hover:text-red-800"
                  >
                    Delete
                  </button>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>

      {showModal && (
        <div className="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50">
          <div className="bg-white rounded-lg p-6 w-full max-w-md">
            <h2 className="text-2xl font-bold mb-4">
              {editMode ? 'Edit Barang' : 'Tambah Barang'}
            </h2>
            <form onSubmit={handleSubmit}>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">Kode Barang</label>
                <input
                  type="text"
                  value={formData.kode_barang}
                  onChange={(e) => setFormData({ ...formData, kode_barang: e.target.value })}
                  className="w-full px-3 py-2 border rounded-md"
                  required
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">Nama Barang</label>
                <input
                  type="text"
                  value={formData.nama_barang}
                  onChange={(e) => setFormData({ ...formData, nama_barang: e.target.value })}
                  className="w-full px-3 py-2 border rounded-md"
                  required
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">Deskripsi</label>
                <textarea
                  value={formData.deskripsi}
                  onChange={(e) => setFormData({ ...formData, deskripsi: e.target.value })}
                  className="w-full px-3 py-2 border rounded-md"
                  rows="3"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">Satuan</label>
                <input
                  type="text"
                  value={formData.satuan}
                  onChange={(e) => setFormData({ ...formData, satuan: e.target.value })}
                  className="w-full px-3 py-2 border rounded-md"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">Harga Beli</label>
                <input
                  type="number"
                  value={formData.harga_beli}
                  onChange={(e) => setFormData({ ...formData, harga_beli: e.target.value })}
                  className="w-full px-3 py-2 border rounded-md"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">Harga Jual</label>
                <input
                  type="number"
                  value={formData.harga_jual}
                  onChange={(e) => setFormData({ ...formData, harga_jual: e.target.value })}
                  className="w-full px-3 py-2 border rounded-md"
                />
              </div>
              <div className="mb-4">
                <label className="block text-sm font-medium mb-1">Stok Awal</label>
                <input
                  type="number"
                  value={formData.stok_akhir}
                  onChange={(e) => setFormData({ ...formData, stok_akhir: e.target.value })}
                  className="w-full px-3 py-2 border rounded-md"
                  placeholder="0"
                />
              </div>
              <div className="flex justify-end gap-2">
                <button
                  type="button"
                  onClick={handleCloseModal}
                  className="px-4 py-2 border rounded-md hover:bg-gray-50"
                >
                  Cancel
                </button>
                <button
                  type="submit"
                  className="px-4 py-2 bg-blue-600 text-white rounded-md hover:bg-blue-700"
                >
                  {editMode ? 'Update' : 'Create'}
                </button>
              </div>
            </form>
          </div>
        </div>
      )}
    </div>
  );
}

export default MasterBarang;
