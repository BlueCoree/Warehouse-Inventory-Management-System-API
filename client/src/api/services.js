import api from './axios';

export const authService = {
  login: (credentials) => api.post('/auth/login', credentials),
  register: (userData) => api.post('/auth/register', userData),
};

export const barangService = {
  getAll: (params) => api.get('/barang', { params }),
  getWithStok: () => api.get('/barang/stok'),
  getById: (id) => api.get(`/barang/${id}`),
  create: (data) => api.post('/barang', data),
  update: (id, data) => api.put(`/barang/${id}`, data),
  delete: (id) => api.delete(`/barang/${id}`),
};

export const stokService = {
  getAll: () => api.get('/stok'),
  getByBarang: (barangId) => api.get(`/stok/${barangId}`),
  getHistory: () => api.get('/history-stok'),
  getHistoryByBarang: (barangId) => api.get(`/history-stok/${barangId}`),
};

export const pembelianService = {
  create: (data) => api.post('/pembelian', data),
  getAll: () => api.get('/pembelian'),
  getById: (id) => api.get(`/pembelian/${id}`),
  selesaikan: (id) => api.put(`/pembelian/${id}/selesaikan`),
  cancel: (id) => api.put(`/pembelian/${id}/cancel`),
};

export const penjualanService = {
  create: (data) => api.post('/penjualan', data),
  getAll: () => api.get('/penjualan'),
  getById: (id) => api.get(`/penjualan/${id}`),
};

export const laporanService = {
  stok: () => api.get('/laporan/stok'),
  penjualan: (params) => api.get('/laporan/penjualan', { params }),
  pembelian: (params) => api.get('/laporan/pembelian', { params }),
};
