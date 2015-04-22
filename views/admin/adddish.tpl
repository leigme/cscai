{{define "adddish"}}
<div class="modal fade" id="addModal" tabindex="-1" role="dialog" aria-labelledby="myModalLabel" aria-hidden="true">
  <div class="modal-dialog modal-lg">
    <div class="modal-content">
      <form enctype="multipart/form-data" action="/admin/controller/dish" method="POST">
        <div class="modal-header">
          <button type="button" class="close" data-dismiss="modal" aria-label="Close"><span aria-hidden="true">&times;</span></button>
          <h4 class="modal-title" id="myModalLabel">添加菜品</h4>
        </div>
        <div class="modal-body">
            <div class="container">
                <div class="row">
                    <div class="col-md-4">
                        <div class="form-group">
                            <input type="text" class="form-control" name="name" placeholder="菜品名称">
                        </div>
                        <div class="form-group">
                            <input type="text" class="form-control" name="price" placeholder="参考价格">
                        </div>
                        <div class="form-group">
                            <label for="picurl">上传图片</label>
                                <input type="file" name="picurl">
                        </div>
                    </div>
                    <div class="col-md-4">
                        <div class="form-group">
                            <div class="input-group">
                                <input type="text" class="form-control" id="datetimepicker" name="updatedtime" aria-describedby="basic-addon1" readonly>
                                <span class="input-group-addon" id="basic-addon1"><i class="icon-th"></i></span>
                            </div>
                        </div>
                        <div class="form-group">
                            <div class="input-group">
                              <select name="status">
                                    <option value="true">选择状态</option>
                                    <option value="true">显示</option>
                                    <option value="false">隐藏</option>
                              </select>
                            </div>
                        </div>
                    </div>
                </div>

                <div class="row">
                    <div class="col-md-8">
                      <label for="content">简 介</label>
                      {{template "editor" .}}
                      <textarea id="texteditor" name="synopsis" style="visibility:hidden"></textarea>
                    </div>
                </div>
                
            </div>
        </div>
        <div class="modal-footer">
          <button type="submit" id="adddish" class="btn btn-primary">提 交</button>
          <button type="button" class="btn btn-default" data-dismiss="modal">取 消</button>
        </div>
      </form>
    </div>
  </div>
</div>
</div>
<script type="text/javascript">
$('#datetimepicker').datetimepicker({
    fromat:'yyyy-mm-dd',
    minView:'month',
    language:'zh-CN',
    todayBtn:'linked',
    autoclose:'true',
});
$('#adddish').click(function(event) {
  var content = $('#editor').html();
  $('#texteditor').val(content);
});
</script>
{{end}}