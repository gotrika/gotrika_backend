<!-- load GOtrika tracker -->
<script type="text/javascript" src="{{ .url }}/static/tracker.v1.min.js"></script>
<!-- init GOtrika tracker -->
<script>
    GOtrika.Init('{{ .code }}', "{{ .url }}/collect/")
</script>