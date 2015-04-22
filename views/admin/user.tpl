{{define "user"}}
<thead>
  <tr>
    <th>#</th>
    <th>名称</th>
    <th>创建时间</th>
    <th>修改时间</th>
    <th>操作</th>
  </tr>
</thead>
<tbody>
{{range .User}}
  <tr>
    <td>{{.Id}}</td>
    <td>{{.Name}}</td>
    <td>{{.Createdtime}}</td>
    <td>{{.Updatedtime}}</td>
    <td><a type="button" class="btn btn-primary btn-xs" href="/admin/controller?op=moduser&id={{.Id}}">修改</a> / <a type="button" class="btn btn-danger btn-xs" href="/admin/controller?op=deluser&id={{.Id}}">删除</a></td>
  </tr>
  {{end}}
</tbody>
{{end}}