package ste

type Buffer struct {
    rows []dataRow
    dirty bool
    id int
}

type dataRow struct {
    rowIdex int
    data []byte
    size int
}

func (b *Buffer) NewRow() {
    b.rows = append(b.rows, dataRow{})
}

func (b *Buffer) Insert(x, y int, data []byte) {
    row := b.rows[x].data
    row = row[:len(data)]
    copy(row[(x + len(data)):], row[x:])
    for i := 0; i < len(data); i++ {
        row[x + i] = data[i]
    }
}
