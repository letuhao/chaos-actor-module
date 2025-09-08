# 06 — Hướng dẫn Implement (Step-by-step)

1) **Dimension catalog**: thêm tất cả primary/derived vào registry của core (clamp_default hợp lý).
2) **Config**: khai báo **bảng cảnh giới** → mỗi tiểu cảnh giới có **caps range** cho primary trọng yếu.
3) **Cultivation loop**: nhận **điểm tu luyện/đan dược/pháp môn** → phát **Primary Contributions** (FLAT/MULT).
4) **Derived compute**: áp công thức từ **components/** theo **Combiner rules** (pipeline) → emit Derived.
5) **Caps emit**: theo cảnh giới hiện tại (scope=REALM); nếu vào độ kiếp → thêm EVENT caps.
6) **Output**: gói `SubsystemOutput` trả về Aggregator.
7) **Tests**: golden vectors: (a) Trúc Cơ→Kim Đan pass/fail, (b) Kim Đan Trung→Thượng nâng `qi_purity`, (c) WORLD cấm vực giảm `divine_sense`.

**Nguyên tắc**: hệ **Kim Đan** không tăng **strength/vitality**; các tăng trưởng thân thể thuộc hệ **Luyện Thể**.
