import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { API } from '../App';
import { Card, CardContent, CardHeader, CardTitle } from '../components/ui/card';
import { Users, Briefcase, Target, Calculator } from 'lucide-react';
import { Link } from 'react-router-dom';

const Dashboard = () => {
  const [stats, setStats] = useState({
    jabatan: 0,
    aspek: 0,
    kriteria: 0,
    tenagaKerja: 0,
  });

  useEffect(() => {
    fetchStats();
  }, []);

  const fetchStats = async () => {
    try {
      const [jabatan, aspek, kriteria, tenagaKerja] = await Promise.all([
        axios.get(`${API}/jabatan`),
        axios.get(`${API}/aspek`),
        axios.get(`${API}/kriteria`),
        axios.get(`${API}/tenaga-kerja`),
      ]);

      setStats({
        jabatan: jabatan.data.length,
        aspek: aspek.data.length,
        kriteria: kriteria.data.length,
        tenagaKerja: tenagaKerja.data.length,
      });
    } catch (error) {
      console.error('Error fetching stats:', error);
    }
  };

  const statCards = [
    {
      title: 'Total Jabatan',
      value: stats.jabatan,
      icon: Briefcase,
      color: 'bg-blue-500',
      link: '/jabatan',
    },
    {
      title: 'Total Aspek',
      value: stats.aspek,
      icon: Target,
      color: 'bg-green-500',
      link: '/aspek',
    },
    {
      title: 'Total Kriteria',
      value: stats.kriteria,
      icon: Calculator,
      color: 'bg-purple-500',
      link: '/kriteria',
    },
    {
      title: 'Total Tenaga Kerja',
      value: stats.tenagaKerja,
      icon: Users,
      color: 'bg-orange-500',
      link: '/tenaga-kerja',
    },
  ];

  return (
    <div>
      <div className="mb-8">
        <h1 className="text-3xl font-bold text-gray-900">Dashboard</h1>
        <p className="text-gray-500 mt-2">Sistem Pendukung Keputusan - KPSG Group</p>
      </div>

      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-4 gap-6 mb-8">
        {statCards.map((stat) => {
          const Icon = stat.icon;
          return (
            <Link key={stat.title} to={stat.link}>
              <Card className="card-hover cursor-pointer">
                <CardHeader className="flex flex-row items-center justify-between pb-2">
                  <CardTitle className="text-sm font-medium text-gray-600">
                    {stat.title}
                  </CardTitle>
                  <div className={`${stat.color} p-3 rounded-lg`}>
                    <Icon className="w-5 h-5 text-white" />
                  </div>
                </CardHeader>
                <CardContent>
                  <div className="text-3xl font-bold text-gray-900">{stat.value}</div>
                </CardContent>
              </Card>
            </Link>
          );
        })}
      </div>

      <div className="grid grid-cols-1 lg:grid-cols-2 gap-6">
        <Card>
          <CardHeader>
            <CardTitle>Tentang Sistem</CardTitle>
          </CardHeader>
          <CardContent>
            <p className="text-gray-600 mb-4">
              Sistem Pendukung Keputusan (SPK) ini menggunakan metode Profile Matching untuk
              membantu dalam penempatan tenaga kerja di KPSG Group.
            </p>
            <ul className="space-y-2 text-sm text-gray-600">
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 rounded-full bg-indigo-500" />
                Perhitungan GAP antara nilai aktual dan target
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 rounded-full bg-indigo-500" />
                Analisis Core Factor (CF) dan Secondary Factor (SF)
              </li>
              <li className="flex items-center gap-2">
                <div className="w-2 h-2 rounded-full bg-indigo-500" />
                Ranking otomatis berdasarkan skor total
              </li>
            </ul>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Panduan Singkat</CardTitle>
          </CardHeader>
          <CardContent>
            <ol className="space-y-3 text-sm text-gray-600">
              <li className="flex gap-3">
                <span className="font-semibold text-indigo-600">1.</span>
                <span>Kelola data master (Jabatan, Aspek, Kriteria)</span>
              </li>
              <li className="flex gap-3">
                <span className="font-semibold text-indigo-600">2.</span>
                <span>Tentukan Target Profile untuk setiap jabatan</span>
              </li>
              <li className="flex gap-3">
                <span className="font-semibold text-indigo-600">3.</span>
                <span>Input data Tenaga Kerja dan nilai mereka</span>
              </li>
              <li className="flex gap-3">
                <span className="font-semibold text-indigo-600">4.</span>
                <span>Jalankan perhitungan dan lihat hasil ranking</span>
              </li>
            </ol>
          </CardContent>
        </Card>
      </div>
    </div>
  );
};

export default Dashboard;
