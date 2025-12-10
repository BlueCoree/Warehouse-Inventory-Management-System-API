# API Documentation - Warehouse Inventory Management System

## Authentication

### Register
**POST** `/api/auth/register`

Request Body:
```json
{
  "username": "string",
  "password": "string",
  "email": "string",
  "full_name": "string"
}
```

Response:
```json
{
  "success": true,
  "message": "User registered successfully",
  "data": null
}
```

### Login
**POST** `/api/auth/login`

Request Body:
```json
{
  "username": "string",
  "password": "string"
}
```

Response:
```json
{
  "success": true,
  "message": "Login successful",
  "data": {
    "token": "jwt_token_here",
    "user": {
      "id": 1,
      "username": "string",
      "email": "string",
      "full_name": "string"
    }
  }
}
```

---

## Master Barang
*Requires Authentication Header: `Authorization: Bearer <token>`*

### Get All Barang
**GET** `/api/barang?search=&page=1&limit=10`

Query Parameters:
- `search` (optional): Search by kode or nama barang
- `page` (optional): Page number (default: 1)
- `limit` (optional): Items per page (default: 10)

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [
    {
      "id": 1,
      "kode_barang": "BRG001",
      "nama_barang": "Mouse Logitech",
      "deskripsi": "Mouse wireless",
      "satuan": "pcs",
      "harga_beli": 150000,
      "harga_jual": 200000,
      "created_at": "2024-01-01T00:00:00Z"
    }
  ],
  "meta": {
    "page": 1,
    "limit": 10,
    "total": 50
  }
}
```

### Get Barang with Stok
**GET** `/api/barang/stok`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [
    {
      "id": 1,
      "kode_barang": "BRG001",
      "nama_barang": "Mouse Logitech",
      "stok_akhir": 25
    }
  ]
}
```

### Get Barang Detail
**GET** `/api/barang/{id}`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "id": 1,
    "kode_barang": "BRG001",
    "nama_barang": "Mouse Logitech",
    "deskripsi": "Mouse wireless",
    "satuan": "pcs",
    "harga_beli": 150000,
    "harga_jual": 200000
  }
}
```

### Create Barang
**POST** `/api/barang`

Request Body:
```json
{
  "kode_barang": "BRG001",
  "nama_barang": "Mouse Logitech",
  "deskripsi": "Mouse wireless",
  "satuan": "pcs",
  "harga_beli": 150000,
  "harga_jual": 200000,
  "stok_akhir": 25
}
```

Response:
```json
{
  "success": true,
  "message": "Barang created successfully",
  "data": {
    "id": 1,
    "kode_barang": "BRG001",
    "nama_barang": "Mouse Logitech",
    "deskripsi": "Mouse wireless",
    "satuan": "pcs",
    "harga_beli": 150000,
    "harga_jual": 200000,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:00:00Z"
  }
}
```

### Update Barang
**PUT** `/api/barang/{id}`

Request Body:
```json
{
  "kode_barang": "BRG001",
  "nama_barang": "Mouse Logitech Updated",
  "deskripsi": "Mouse wireless updated",
  "satuan": "pcs",
  "harga_beli": 160000,
  "harga_jual": 210000,
  "stok_akhir": 30
}
```

Response:
```json
{
  "success": true,
  "message": "Barang updated successfully",
  "data": {
    "id": 1,
    "kode_barang": "BRG001",
    "nama_barang": "Mouse Logitech Updated",
    "deskripsi": "Mouse wireless updated",
    "satuan": "pcs",
    "harga_beli": 160000,
    "harga_jual": 210000,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-02T00:00:00Z"
  }
}
```

### Delete Barang
**DELETE** `/api/barang/{id}`

Response:
```json
{
  "success": true,
  "message": "Barang deleted successfully",
  "data": null
}
```

---

## Stok Management
*Requires Authentication*

### Get All Stok
**GET** `/api/stok`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [
    {
      "id": 1,
      "barang_id": 1,
      "stok_akhir": 25,
      "updated_at": "2024-01-01T00:00:00Z",
      "barang": {
        "kode_barang": "BRG001",
        "nama_barang": "Mouse Logitech"
      }
    }
  ]
}
```

### Get Stok by Barang
**GET** `/api/stok/{barang_id}`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "id": 1,
    "barang_id": 1,
    "stok_akhir": 25,
    "updated_at": "2024-01-01T00:00:00Z",
    "barang": {
      "id": 1,
      "kode_barang": "BRG001",
      "nama_barang": "Mouse Logitech",
      "deskripsi": "Mouse wireless",
      "satuan": "pcs",
      "harga_beli": 150000,
      "harga_jual": 200000
    }
  }
}
```

---

## History Stok
*Requires Authentication*

### Get All History
**GET** `/api/history-stok`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "barang_id": 1,
        "nama_barang": "Mouse Logitech",
        "user_id": 1,
        "jenis_transaksi": "masuk",
        "jumlah": 10,
        "stok_sebelum": 15,
        "stok_sesudah": 25,
        "keterangan": "Pembelian BLI001",
        "created_at": "2024-01-01T00:00:00Z",
        "barang": {
          "kode_barang": "BRG001",
          "nama_barang": "Mouse Logitech"
        },
        "user": {
          "username": "admin",
          "full_name": "Administrator"
        }
      }
    ],
    "total": 1
  }
}
```

### Get History by Barang
**GET** `/api/history-stok/{barang_id}`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "barang_id": 1,
        "nama_barang": "Mouse Logitech",
        "user_id": 1,
        "jenis_transaksi": "masuk",
        "jumlah": 10,
        "stok_sebelum": 15,
        "stok_sesudah": 25,
        "keterangan": "Pembelian BLI001",
        "created_at": "2024-01-01T00:00:00Z",
        "barang": {
          "kode_barang": "BRG001",
          "nama_barang": "Mouse Logitech"
        },
        "user": {
          "username": "admin",
          "full_name": "Administrator"
        }
      },
      {
        "id": 2,
        "barang_id": 1,
        "nama_barang": "Mouse Logitech",
        "user_id": 1,
        "jenis_transaksi": "keluar",
        "jumlah": 5,
        "stok_sebelum": 25,
        "stok_sesudah": 20,
        "keterangan": "Penjualan JL001",
        "created_at": "2024-01-02T00:00:00Z",
        "barang": {
          "kode_barang": "BRG001",
          "nama_barang": "Mouse Logitech"
        },
        "user": {
          "username": "admin",
          "full_name": "Administrator"
        }
      }
    ],
    "total": 2
  }
}
```

---

## Pembelian (Purchase)
*Requires Authentication*

### Create Pembelian
**POST** `/api/pembelian`

Request Body:
```json
{
  "no_faktur": "BLI001",
  "supplier": "PT. Supplier ABC",
  "status": "pending",
  "details": [
    {
      "barang_id": 1,
      "qty": 10,
      "harga": 150000
    }
  ]
}
```

Response:
```json
{
  "success": true,
  "message": "Pembelian created successfully",
  "data": {
    "id": 1,
    "no_faktur": "BLI001",
    "supplier": "PT. Supplier ABC",
    "total": 1500000,
    "status": "pending",
    "user_id": 1,
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### Get All Pembelian
**GET** `/api/pembelian`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [
    {
      "id": 1,
      "no_faktur": "BLI001",
      "supplier": "PT. Supplier ABC",
      "total": 1500000,
      "status": "pending",
      "created_at": "2024-01-01T00:00:00Z",
      "user": {
        "username": "admin"
      },
      "details": [
        {
          "id": 1,
          "barang_id": 1,
          "nama_barang": "Mouse Logitech",
          "qty": 10,
          "harga": 150000,
          "subtotal": 1500000
        }
      ]
    }
  ]
}
```

### Get Pembelian Detail
**GET** `/api/pembelian/{id}`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "header": {
      "id": 1,
      "no_faktur": "BLI001",
      "supplier": "PT. Supplier ABC",
      "total": 1500000,
      "status": "pending",
      "user_id": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z",
      "user": {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "full_name": "Administrator"
      }
    },
    "details": [
      {
        "id": 1,
        "barang_id": 1,
        "nama_barang": "Mouse Logitech",
        "qty": 10,
        "harga": 150000,
        "subtotal": 1500000,
        "barang": {
          "kode_barang": "BRG001",
          "nama_barang": "Mouse Logitech",
          "satuan": "pcs"
        }
      }
    ]
  }
}
```

### Selesaikan Pembelian (Complete Purchase)
**PUT** `/api/pembelian/{id}/selesaikan`

*Updates stock when purchase is completed. Only the user who created the purchase can complete it.*

Response:
```json
{
  "success": true,
  "message": "Pembelian completed successfully",
  "data": {
    "id": 1,
    "no_faktur": "BLI001",
    "supplier": "PT. Supplier ABC",
    "total": 1500000,
    "status": "selesai",
    "user_id": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:10:00Z"
  }
}
```

### Cancel Pembelian
**PUT** `/api/pembelian/{id}/cancel`

*Cancels purchase (auto-triggered after 5 minutes timeout)*

Response:
```json
{
  "success": true,
  "message": "Pembelian cancelled successfully",
  "data": {
    "id": 1,
    "no_faktur": "BLI001",
    "supplier": "PT. Supplier ABC",
    "total": 1500000,
    "status": "cancel",
    "user_id": 1,
    "created_at": "2024-01-01T00:00:00Z",
    "updated_at": "2024-01-01T00:05:00Z"
  }
}
```

---

## Penjualan (Sales)
*Requires Authentication*

### Create Penjualan
**POST** `/api/penjualan`

Request Body:
```json
{
  "no_faktur": "JL001",
  "customer": "John Doe",
  "status": "selesai",
  "details": [
    {
      "barang_id": 1,
      "qty": 5,
      "harga": 200000
    }
  ]
}
```

Response:
```json
{
  "success": true,
  "message": "Penjualan created successfully",
  "data": {
    "id": 1,
    "no_faktur": "JL001",
    "customer": "John Doe",
    "total": 1000000,
    "status": "selesai",
    "created_at": "2024-01-01T00:00:00Z"
  }
}
```

### Get All Penjualan
**GET** `/api/penjualan`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": [
    {
      "id": 1,
      "no_faktur": "JL001",
      "customer": "John Doe",
      "total": 1000000,
      "status": "selesai",
      "created_at": "2024-01-01T00:00:00Z",
      "user": {
        "username": "admin"
      },
      "details": [
        {
          "id": 1,
          "barang_id": 1,
          "nama_barang": "Mouse Logitech",
          "qty": 5,
          "harga": 200000,
          "subtotal": 1000000
        }
      ]
    },
    {
      "id": 2,
      "no_faktur": "JL002",
      "customer": "Jane Smith",
      "total": 750000,
      "status": "selesai",
      "created_at": "2024-01-02T00:00:00Z",
      "user": {
        "username": "admin"
      },
      "details": [
        {
          "id": 2,
          "barang_id": 2,
          "nama_barang": "Keyboard Mechanical",
          "qty": 3,
          "harga": 250000,
          "subtotal": 750000
        }
      ]
    }
  ]
}
```

### Get Penjualan Detail
**GET** `/api/penjualan/{id}`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "header": {
      "id": 1,
      "no_faktur": "JL001",
      "customer": "John Doe",
      "total": 1000000,
      "status": "selesai",
      "user_id": 1,
      "created_at": "2024-01-01T00:00:00Z",
      "updated_at": "2024-01-01T00:00:00Z",
      "user": {
        "id": 1,
        "username": "admin",
        "email": "admin@example.com",
        "full_name": "Administrator"
      }
    },
    "details": [
      {
        "id": 1,
        "barang_id": 1,
        "nama_barang": "Mouse Logitech",
        "qty": 5,
        "harga": 200000,
        "subtotal": 1000000,
        "barang": {
          "kode_barang": "BRG001",
          "nama_barang": "Mouse Logitech",
          "satuan": "pcs"
        }
      }
    ]
  }
}
```

---

## Laporan (Reports)
*Requires Authentication*

### Laporan Stok Akhir
**GET** `/api/laporan/stok`

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "barang_id": 1,
        "kode_barang": "BRG001",
        "nama_barang": "Mouse Logitech",
        "satuan": "pcs",
        "stok_akhir": 25,
        "harga_beli": 150000,
        "harga_jual": 200000,
        "nilai_stok": 3750000,
        "deleted": false
      }
    ],
    "total_items": 10,
    "total_nilai": 15000000
  }
}
```

### Laporan Penjualan
**GET** `/api/laporan/penjualan?start_date=2024-01-01&end_date=2024-12-31`

Query Parameters:
- `start_date` (optional): Filter start date (YYYY-MM-DD)
- `end_date` (optional): Filter end date (YYYY-MM-DD)

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "no_faktur": "JL001",
        "customer": "John Doe",
        "total": 1000000,
        "status": "selesai",
        "created_at": "2024-01-01T00:00:00Z",
        "username": "admin",
        "total_items": 2,
        "details": [
          {
            "nama_barang": "Mouse Logitech",
            "qty": 5,
            "harga": 200000,
            "subtotal": 1000000
          }
        ]
      }
    ],
    "total_transaksi": 1,
    "grand_total": 1000000
  }
}
```

### Laporan Pembelian
**GET** `/api/laporan/pembelian?start_date=2024-01-01&end_date=2024-12-31`

Query Parameters:
- `start_date` (optional): Filter start date (YYYY-MM-DD)
- `end_date` (optional): Filter end date (YYYY-MM-DD)

Response:
```json
{
  "success": true,
  "message": "Data retrieved successfully",
  "data": {
    "data": [
      {
        "id": 1,
        "no_faktur": "BLI001",
        "supplier": "PT. Supplier ABC",
        "total": 1500000,
        "status": "selesai",
        "created_at": "2024-01-01T00:00:00Z",
        "username": "admin",
        "total_items": 1,
        "details": [
          {
            "nama_barang": "Mouse Logitech",
            "qty": 10,
            "harga": 150000,
            "subtotal": 1500000
          }
        ]
      },
      {
        "id": 2,
        "no_faktur": "BLI002",
        "supplier": "PT. Supplier XYZ",
        "total": 2000000,
        "status": "selesai",
        "created_at": "2024-01-05T00:00:00Z",
        "username": "admin",
        "total_items": 2,
        "details": [
          {
            "nama_barang": "Keyboard Mechanical",
            "qty": 5,
            "harga": 250000,
            "subtotal": 1250000
          },
          {
            "nama_barang": "Headset Gaming",
            "qty": 3,
            "harga": 250000,
            "subtotal": 750000
          }
        ]
      }
    ],
    "total_transaksi": 2,
    "grand_total": 3500000
  }
}
```

---

## Error Responses

All endpoints may return error responses in this format:

```json
{
  "success": false,
  "message": "Error message here",
  "data": null
}
```

Common HTTP Status Codes:
- `200` - Success
- `400` - Bad Request
- `401` - Unauthorized
- `404` - Not Found
- `422` - Unprocessable Entity (validation error)
- `500` - Internal Server Error

---

