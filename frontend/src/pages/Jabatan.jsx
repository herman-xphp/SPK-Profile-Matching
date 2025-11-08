import React, { useState, useEffect, useContext } from 'react';
import axios from 'axios';
import { API, AuthContext } from '../App';
import { Button } from '../components/ui/button';
import { Input } from '../components/ui/input';
import { Label } from '../components/ui/label';
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card';
import { Dialog, DialogContent, DialogHeader, DialogTitle, DialogTrigger } from '../components/ui/dialog';
import { toast } from 'sonner';
import { Plus, Edit, Trash2 } from 'lucide-react';

const Jabatan = () => {
  const { user } = useContext(AuthContext);
  const [data, setData] = useState([]);
  const [open, setOpen] = useState(false);
  const [editMode, setEditMode] = useState(false);
  const [formData, setFormData] = useState({ id: '', nama: '', deskripsi: '' });

  useEffect(() => {
    fetchData();
  }, []);

  const fetchData = async () => {
    try {
      const response = await axios.get(`${API}/jabatan`);
      setData(response.data);
    } catch (error) {
      toast.error('Gagal memuat data jabatan');
    }
  };

  const handleSubmit = async (e) => {
    e.preventDefault();
    try {
      if (editMode) {
        await axios.put(`${API}/jabatan/${formData.id}`, { nama: formData.nama, deskripsi: formData.deskripsi });
        toast.success('Jabatan berhasil diperbarui');
      } else {
        await axios.post(`${API}/jabatan`, { nama: formData.nama, deskripsi: formData.deskripsi });
        toast.success('Jabatan berhasil ditambahkan');
      }
      setOpen(false);
      resetForm();
      fetchData();
    } catch (error) {
      toast.error(error.response?.data?.detail || 'Operasi gagal');
    }
  };

  const handleDelete = async (id) => {
    if (!window.confirm('Yakin ingin menghapus jabatan ini?')) return;
    try {
      await axios.delete(`${API}/jabatan/${id}`);
      toast.success('Jabatan berhasil dihapus');
      fetchData();
    } catch (error) {
      toast.error('Gagal menghapus jabatan');
    }
  };

  const resetForm = () => {
    setFormData({ id: '', nama: '', deskripsi: '' });
    setEditMode(false);
  };

  const handleEdit = (item) => {
    setFormData(item);
    setEditMode(true);
    setOpen(true);
  };

  const handleAdd = () => {
    resetForm();
    setOpen(true);
  };

  const isAdmin = user?.role === 'admin';

  return (
    <div>
      <div className="flex justify-between items-center mb-6">
        <div>
          <h1 className="text-2xl font-bold text-gray-900">Data Jabatan</h1>
          <p className="text-gray-500 mt-1">Kelola data jabatan di KPSG Group</p>
        </div>
        {isAdmin && (
          <Dialog open={open} onOpenChange={setOpen}>
            <DialogTrigger asChild>
              <Button onClick={handleAdd} data-testid="add-jabatan-button">
                <Plus size={16} className="mr-2" />
                Tambah Jabatan
              </Button>
            </DialogTrigger>
            <DialogContent>
              <DialogHeader>
                <DialogTitle>{editMode ? 'Edit Jabatan' : 'Tambah Jabatan'}</DialogTitle>
              </DialogHeader>
              <form onSubmit={handleSubmit} className="space-y-4">
                <div className="space-y-2">
                  <Label htmlFor="nama">Nama Jabatan</Label>
                  <Input
                    id="nama"
                    data-testid="jabatan-nama-input"
                    value={formData.nama}
                    onChange={(e) => setFormData({ ...formData, nama: e.target.value })}
                    required
                  />
                </div>
                <div className="space-y-2">
                  <Label htmlFor="deskripsi">Deskripsi</Label>
                  <Input
                    id="deskripsi"
                    data-testid="jabatan-deskripsi-input"
                    value={formData.deskripsi || ''}
                    onChange={(e) => setFormData({ ...formData, deskripsi: e.target.value })}
                  />
                </div>
                <div className="flex gap-2">
                  <Button type="submit" data-testid="jabatan-submit-button" className="flex-1">
                    {editMode ? 'Update' : 'Simpan'}
                  </Button>
                  <Button
                    type="button"
                    variant="outline"
                    onClick={() => setOpen(false)}
                    className="flex-1"
                  >
                    Batal
                  </Button>
                </div>
              </form>
            </DialogContent>
          </Dialog>
        )}
      </div>

      <Card>
        <CardHeader>
          <CardTitle>Daftar Jabatan</CardTitle>
        </CardHeader>
        <CardContent>
          {data.length === 0 ? (
            <p className="text-center text-gray-500 py-8">Belum ada data jabatan</p>
          ) : (
            <div className="overflow-x-auto">
              <table className="w-full">
                <thead>
                  <tr className="border-b">
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">No</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Nama Jabatan</th>
                    <th className="text-left py-3 px-4 font-semibold text-gray-700">Deskripsi</th>
                    {isAdmin && <th className="text-center py-3 px-4 font-semibold text-gray-700">Aksi</th>}
                  </tr>
                </thead>
                <tbody>
                  {data.map((item, index) => (
                    <tr key={item.id} className="border-b hover:bg-gray-50" data-testid={`jabatan-row-${index}`}>
                      <td className="py-3 px-4">{index + 1}</td>
                      <td className="py-3 px-4 font-medium">{item.nama}</td>
                      <td className="py-3 px-4 text-gray-600">{item.deskripsi || '-'}</td>
                      {isAdmin && (
                        <td className="py-3 px-4">
                          <div className="flex gap-2 justify-center">
                            <Button
                              size="sm"
                              variant="outline"
                              onClick={() => handleEdit(item)}
                              data-testid={`edit-jabatan-${index}`}
                            >
                              <Edit size={14} />
                            </Button>
                            <Button
                              size="sm"
                              variant="destructive"
                              onClick={() => handleDelete(item.id)}
                              data-testid={`delete-jabatan-${index}`}
                            >
                              <Trash2 size={14} />
                            </Button>
                          </div>
                        </td>
                      )}
                    </tr>
                  ))}
                </tbody>
              </table>
            </div>
          )}
        </CardContent>
      </Card>
    </div>
  );
};

export default Jabatan;
