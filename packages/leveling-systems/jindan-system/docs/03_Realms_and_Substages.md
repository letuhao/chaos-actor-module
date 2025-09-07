# 03 — Cảnh Giới & Tiểu Cảnh Giới (dựa cảm hứng Phong Thần)

## Đại Cảnh Giới (gợi ý)
0. **Phàm** (Mortal)  
1. **Luyện Khí** (Qi Condensation) — **Tầng 1..9**  
2. **Trúc Cơ** (Foundation Establishment) — **Sơ / Trung / Hậu / Viên Mãn**  
3. **Kim Đan** (Golden Core) — **Hạ / Trung / Thượng / Viên Mãn** + **Phẩm Đan** (Hạ / Trung / Thượng / Cực)  
4. **Nguyên Anh** (Nascent Soul) — **Tiểu / Trung / Đại / Viên Mãn**  
5. **Hóa Thần** (Spirit Transformation) — **Sơ / Trung / Hậu / Viên Mãn**  
6. **Luyện Hư** (Void Refining)  
7. **Hợp Thể** (Integration)  
8. **Đại Thừa** (Great Ascension)  
9. **Phi Thăng / Tán Tiên → Địa/Thiên/Kim/Thái Ất** (tuỳ thế giới quan)

> Tuỳ game, có thể dừng ở **Kim Đan/Nguyên Anh** nếu muốn chiều sâu thay vì dài.

## Ngưỡng/Caps theo Cảnh Giới (scope=REALM)
- Mỗi **tiểu cảnh giới** định nghĩa **khoảng cap** cho **Primary Stats** then chốt.
- Ví dụ (trích):
```
REALM: 'Kim Đan: Thượng Phẩm'
caps:
  dantian_capacity:        [ 50_000 .. 120_000 ]
  dantian_compression:     [ 0.50   .. 1.20   ]
  dantian_stability:       [ 120    .. 300    ]
  meridian_conductivity:   [ 500    .. 1_600  ]
  meridian_toughness:      [ 300    .. 800    ]
  qi_purity:               [ 0.40   .. 0.85   ]
  shen_depth:              [ 200    .. 700    ]
  shen_clarity:            [ 150    .. 600    ]
  shen_control:            [ 120    .. 500    ]
  eff_qi_to_shen:          [ 0.10   .. 0.40   ]
  eff_jing_to_qi:          [ 0.15   .. 0.50   ]
```

## Sự kiện độ kiếp (scope=EVENT)
- Các bậc chuyển cảnh (Trúc Cơ→Kim Đan, Kim Đan→Nguyên Anh…) sinh **caps EVENT**:  
  - `tribulation_resist` **HARD_MIN**: yêu cầu tối thiểu để vượt kiếp.  
  - Kẹp `qi_purity`, `dantian_stability` tối thiểu.  
  - **Nếu fail**: emit debuff caps tạm thời và giảm Primary theo tỉ lệ.
